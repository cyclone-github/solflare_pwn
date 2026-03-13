package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/storage"
	"github.com/syndtr/goleveldb/leveldb/table"
)

/*
Cyclone's Solflare Vault Extractor
https://github.com/cyclone-github/solflare_pwn
POC tool to extract Solflare vault wallets
This tool is proudly the first Solflare Vault Extractor
coded by cyclone in Go

GNU General Public License v2.0
https://github.com/cyclone-github/solflare_pwn/blob/main/LICENSE

version history
v0.1.0-2025-02-12;
	initial release
v0.2.0-2025-03-13;
	added walletdata support (latest Chrome extension)
	added dump file input (key: value format)
v0.3.1; 2026-03-13;
	add support for latest extension
*/

// version func
func versionFunc() {
	fmt.Fprintln(os.Stderr, "Cyclone's Solflare Vault Extractor v0.3.1; 2026-03-13")
	fmt.Fprintln(os.Stderr, "https://github.com/cyclone-github/solflare_pwn")
}

// help func
func helpFunc() {
	versionFunc()
	str := `Example Usage:
./solflare_extractor.bin [-version] [-help] [solflare_vault_dir | dump_file]
./solflare_extractor.bin bhhhlbepdkbapadjdnnojkbgioiodbic/
./solflare_extractor.bin new_wallet_dump.txt

Supports:
- LevelDB directory (Chrome Local Extension Settings)
- Dump file (key: value format, one entry per line, e.g. walletdata: {...})

Default Solflare vault locations for Chrome extensions:

Linux:
/home/$USER/.config/google-chrome/Default/Local\ Extension\ Settings/bhhhlbepdkbapadjdnnojkbgioiodbic/

Mac:
Library>Application Support>Google>Chrome>Default>Local Extension Settings>bhhhlbepdkbapadjdnnojkbgioiodbic

Windows:
C:\Users\$USER\AppData\Local\Google\Chrome\User Data\Default\Local Extension Settings\bhhhlbepdkbapadjdnnojkbgioiodbic`
	fmt.Fprintln(os.Stderr, str)
}

// print welcome screen
func printWelcomeScreen() {
	fmt.Println(" ----------------------------------------------------- ")
	fmt.Println("|        Cyclone's Solflare Vault Hash Extractor       |")
	fmt.Println("|        Use Solflare Vault Decryptor to decrypt       |")
	fmt.Println("|    https://github.com/cyclone-github/solflare_pwn    |")
	fmt.Println(" ----------------------------------------------------- ")
}

// solflare vault
type SolflareVault struct {
	Data struct {
		Digest     string `json:"digest"`
		Encoding   string `json:"encoding"`
		Encrypted  string `json:"encrypted64"`
		Iterations int    `json:"iterations"`
		Kdf        string `json:"kdf"`
		Nonce      string `json:"nonce64"`
		Salt       string `json:"salt64"`
	} `json:"data"`
	Locked bool `json:"locked"`
}

// processLevelDB
func processLevelDB(key, value []byte) {
	keyStr := string(key)
	valStr := string(value)

	// extract encrypted mnemonic seed phrase (solflaredata=old, walletdata=latest extension)
	if strings.Contains(keyStr, "solflaredata") || strings.Contains(keyStr, "walletdata") || strings.Contains(valStr, "\"solflaredata\"") {
		var vault SolflareVault
		err := json.Unmarshal(value, &vault)
		if err != nil {
			fmt.Println("Error parsing Solflare vault:", err)
			return
		}
		decodedSalt, _ := base64.StdEncoding.DecodeString(vault.Data.Salt)
		decodedNonce, _ := base64.StdEncoding.DecodeString(vault.Data.Nonce)
		decodedEncrypted, _ := base64.StdEncoding.DecodeString(vault.Data.Encrypted)

		fmt.Println("\nEncrypted Solflare Vault:")
		// $solflare${iterations}${salt}${nonce}${encrypted}
		fmt.Printf("$solflare$%d$%x$%x$%x", vault.Data.Iterations, decodedSalt, decodedNonce, decodedEncrypted)
	}

	// extract password
	if strings.Contains(keyStr, "solflarexpass") {
		var password string
		err := json.Unmarshal(value, &password)
		if err != nil {
			fmt.Println("Error parsing password:", err)
			return
		}
		decodedPassword, err := base64.StdEncoding.DecodeString(password)
		if err != nil {
			fmt.Println("Error decoding base64 password:", err)
			return
		}

		fmt.Printf("$%x\n", decodedPassword)

	}
}

func dumpRawLDBFiles(dirPath string) error {
	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Failed to access path %s: %v", path, err)
			return nil
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".ldb") {
			err = dumpRawLDBFile(path)
			if err != nil {
				log.Printf("Failed to dump file %s: %v", path, err)
			}
		}
		return nil
	})
}

func dumpRawLDBFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	reader, err := table.NewReader(file, fileInfo.Size(), storage.FileDesc{Type: storage.TypeTable, Num: 0}, nil, nil, &opt.Options{})
	if err != nil {
		return fmt.Errorf("failed to create table reader: %w", err)
	}
	defer reader.Release()

	iter := reader.NewIterator(nil, nil)
	defer iter.Release()

	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		processLevelDB(key, filterPrintableBytes(value))
	}
	if err := iter.Error(); err != nil {
		return fmt.Errorf("iterator error: %w", err)
	}

	return nil
}

func filterPrintableBytes(data []byte) []byte {
	printable := make([]rune, 0, len(data))
	for _, b := range data {
		if unicode.IsPrint(rune(b)) {
			printable = append(printable, rune(b))
		} else {
			printable = append(printable, '.')
		}
	}
	return []byte(string(printable))
}

// processDumpFile processes a dump file with "key: value" format (one entry per line)
func processDumpFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open dump file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Increase buffer for long lines (walletdata JSON can be large)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		idx := strings.Index(line, ": ")
		if idx < 0 {
			continue
		}
		key := strings.TrimSpace(line[:idx])
		value := strings.TrimSpace(line[idx+2:])
		processLevelDB([]byte(key), []byte(value))
	}
	return scanner.Err()
}

// main
func main() {
	cycloneFlag := flag.Bool("cyclone", false, "")
	versionFlag := flag.Bool("version", false, "Program version")
	helpFlag := flag.Bool("help", false, "Program usage instructions")
	flag.Parse()

	// run sanity checks for special flags
	if *versionFlag {
		versionFunc()
		os.Exit(0)
	}
	if *cycloneFlag {
		line := "Q29kZWQgYnkgY3ljbG9uZSA7KQo="
		str, _ := base64.StdEncoding.DecodeString(line)
		fmt.Println(string(str))
		os.Exit(0)
	}
	if *helpFlag {
		helpFunc()
		os.Exit(0)
	}

	inputPath := flag.Arg(0)
	if inputPath == "" {
		fmt.Fprintln(os.Stderr, "Error: Solflare vault directory or dump file is required")
		helpFunc()
		os.Exit(1)
	}

	printWelcomeScreen()

	info, err := os.Stat(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error accessing path %s: %v\n", inputPath, err)
		os.Exit(1)
	}

	// If input is a file, process as dump file (key: value format)
	if !info.IsDir() {
		if err := processDumpFile(inputPath); err != nil {
			fmt.Fprintf(os.Stderr, "Error processing dump file: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	// Input is directory - use LevelDB
	db, err := leveldb.OpenFile(inputPath, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening Vault:", err)
		fmt.Println("Attempting to dump raw .ldb files...")
		err = dumpRawLDBFiles(inputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to dump raw .ldb files: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}
	defer db.Close()

	iter := db.NewIterator(nil, nil)
	defer iter.Release()

	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		processLevelDB(key, value)
	}
}
