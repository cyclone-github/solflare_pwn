package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/crypto/pbkdf2"
)

// vault struct for Solflare Wallet Vaults
type Vault struct {
	EncryptedData   []byte
	Salt            []byte
	Nonce           []byte
	Iterations      int
	Decrypted       int32
	Kdf             string
	VaultText       string
	DecodedPassword []byte
}

// isValid function as placeholder, always returning true
func isValid(s []byte) bool {
	return true
}

// decryptVault w/secretbox using pbkdf2 KDF
func decryptVault(encryptedData, password, salt, nonce []byte, iterations int, kdf string) ([]byte, string, error) {
	if len(nonce) != 24 {
		return nil, "", fmt.Errorf("nonce must be exactly 24 bytes long")
	}
	if len(salt) < 8 {
		return nil, "", fmt.Errorf("salt must be at least 8 bytes long")
	}

	var key []byte
	switch kdf {
	case "pbkdf2":
		key = pbkdf2.Key(password, salt, iterations, 32, sha256.New)
	default:
		return nil, "", fmt.Errorf("unsupported KDF: %s", kdf)
	}

	var nonceArray [24]byte
	copy(nonceArray[:], nonce)
	var keyArray [32]byte
	copy(keyArray[:], key)

	decrypted, ok := secretbox.Open(nil, encryptedData, &nonceArray, &keyArray)
	if !ok {
		return nil, "", fmt.Errorf("decryption failed")
	}

	var decryptedJSON map[string]interface{}
	if err := json.Unmarshal(decrypted, &decryptedJSON); err != nil {
		return nil, "", fmt.Errorf("error parsing decrypted JSON: %v", err)
	}

	// Standard structure: wallets -> {id: {data: {base: base58_mnemonic}}}
	wallets, ok := decryptedJSON["wallets"].(map[string]interface{})
	if ok {
		var walletID string
		for id := range wallets {
			walletID = id
			break
		}
		if walletID != "" {
			walletData, ok := wallets[walletID].(map[string]interface{})
			if ok {
				data, ok := walletData["data"].(map[string]interface{})
				if ok {
					base, ok := data["base"].(string)
					if ok && base != "" {
						mnemonicBytes := base58.Decode(base)
						seedPhrase := string(mnemonicBytes)
						if len(seedPhrase) > 0 {
							return decrypted, seedPhrase, nil
						}
					}
				}
			}
		}
	}

	// Fallback: search for "base" or "mnemonic" anywhere in nested structure (latest extension may vary)
	if seedPhrase := extractMnemonicFromJSON(decryptedJSON); seedPhrase != "" {
		return decrypted, seedPhrase, nil
	}

	return nil, "", fmt.Errorf("invalid JSON structure: could not extract mnemonic")
}

// extractMnemonicFromJSON recursively searches for base58 mnemonic in decrypted JSON
func extractMnemonicFromJSON(m map[string]interface{}) string {
	for k, v := range m {
		switch val := v.(type) {
		case string:
			if (k == "base" || k == "mnemonic") && len(val) > 10 {
				decoded := base58.Decode(val)
				if len(decoded) > 0 {
					return string(decoded)
				}
			}
		case map[string]interface{}:
			if s := extractMnemonicFromJSON(val); s != "" {
				return s
			}
		}
	}
	return ""
}

// rawWalletData struct for parsing walletdata JSON (latest extension format)
type rawWalletData struct {
	Data struct {
		Encrypted64 string `json:"encrypted64"`
		Iterations  int    `json:"iterations"`
		Nonce64     string `json:"nonce64"`
		Salt64      string `json:"salt64"`
		Kdf         string `json:"kdf"`
	} `json:"data"`
}

// parseRawWalletJSON parses walletdata JSON format and returns a Vault
func parseRawWalletJSON(line string) (*Vault, bool) {
	// Strip optional "walletdata:" prefix
	jsonStr := strings.TrimSpace(line)
	if strings.HasPrefix(jsonStr, "walletdata:") {
		jsonStr = strings.TrimSpace(strings.TrimPrefix(jsonStr, "walletdata:"))
	}
	if !strings.HasPrefix(jsonStr, "{") {
		return nil, false
	}

	var raw rawWalletData
	if err := json.Unmarshal([]byte(jsonStr), &raw); err != nil {
		return nil, false
	}
	if raw.Data.Encrypted64 == "" || raw.Data.Nonce64 == "" || raw.Data.Salt64 == "" {
		return nil, false
	}

	salt, err := base64.StdEncoding.DecodeString(raw.Data.Salt64)
	if err != nil || len(salt) < 8 {
		return nil, false
	}
	nonce, err := base64.StdEncoding.DecodeString(raw.Data.Nonce64)
	if err != nil || len(nonce) != 24 {
		return nil, false
	}
	encrypted, err := base64.StdEncoding.DecodeString(raw.Data.Encrypted64)
	if err != nil || len(encrypted) == 0 {
		return nil, false
	}

	iterations := raw.Data.Iterations
	if iterations <= 0 {
		iterations = 600000
	}
	kdf := raw.Data.Kdf
	if kdf == "" {
		kdf = "pbkdf2"
	}

	return &Vault{
		EncryptedData: encrypted,
		Salt:          salt,
		Nonce:         nonce,
		Iterations:    iterations,
		Kdf:           kdf,
		VaultText:     line,
	}, true
}

// parse Solflare vault
func readVaultData(filePath string) ([]Vault, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var vaults []Vault
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Try raw walletdata JSON first (latest extension)
		if v, ok := parseRawWalletJSON(line); ok {
			vaults = append(vaults, *v)
			continue
		}

		if !strings.HasPrefix(line, "$solflare$") {
			log.Printf("Skipping invalid line: %.80s...\n", line)
			continue
		}

		parts := strings.Split(line, "$")
		if len(parts) < 6 {
			log.Printf("Invalid Solflare hash format: %s\n", line)
			continue
		}

		iterations, err := strconv.Atoi(parts[2])
		if err != nil || iterations <= 0 {
			log.Printf("Invalid iteration count: %s\n", parts[2])
			continue
		}

		salt, err := hex.DecodeString(parts[3])
		if err != nil || len(salt) < 8 {
			log.Printf("Error decoding hex salt: %s\n", parts[3])
			continue
		}

		nonce, err := hex.DecodeString(parts[4])
		if err != nil || len(nonce) != 24 {
			log.Printf("Error decoding hex nonce: %s\n", parts[4])
			continue
		}

		encryptedData, err := hex.DecodeString(parts[5])
		if err != nil || len(encryptedData) == 0 {
			log.Printf("Error decoding hex encrypted data: %s\n", parts[5])
			continue
		}

		var decodedPassword []byte
		if len(parts) > 6 {
			passwordHex := parts[6]
			decodedPassword, err = hex.DecodeString(passwordHex)
			if err != nil || len(decodedPassword) == 0 {
				log.Printf("Error decoding hex password: %s\n", passwordHex)
				continue
			}
			decodedPassword = []byte(base64.StdEncoding.EncodeToString(decodedPassword))
		}

		vault := Vault{
			EncryptedData:   encryptedData,
			Salt:            salt,
			Nonce:           nonce,
			Iterations:      iterations,
			Kdf:             "pbkdf2",
			VaultText:       line,
			DecodedPassword: decodedPassword,
		}
		vaults = append(vaults, vault)
	}

	return vaults, nil
}
