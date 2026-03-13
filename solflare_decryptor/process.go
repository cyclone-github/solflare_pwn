package main

import (
	"bufio"
	"crypto/sha512"
	"fmt"
	"log"
	"os"
	"sync"
	"sync/atomic"

	"github.com/btcsuite/btcutil/base58"
)

// process logic
func startProc(wordlistFileFlag string, outputPath string, numGoroutines int, vaults []Vault, crackedCount *int32, linesProcessed *int32, stopChan chan struct{}) {
	var file *os.File
	var err error

	if wordlistFileFlag == "" {
		file = os.Stdin
	} else {
		file, err = os.Open(wordlistFileFlag)
		if err != nil {
			log.Fatalf("Error opening file: %v\n", err)
		}
		defer file.Close()
	}

	var outputFile *os.File
	if outputPath != "" {
		outputFile, err = os.OpenFile(outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Error opening output file: %v", err)
		}
		defer outputFile.Close()
	}

	var writer *bufio.Writer
	if outputPath != "" {
		writer = bufio.NewWriter(outputFile)
	} else {
		writer = bufio.NewWriter(os.Stdout)
	}
	defer writer.Flush()

	var (
		writerMu sync.Mutex
		wg       sync.WaitGroup
	)

	// start worker goroutines
	linesCh := make(chan []byte, 1000)
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for password := range linesCh {
				processPassword(password, vaults, &writerMu, writer, crackedCount, linesProcessed, stopChan)
			}
		}()
	}

	// read lines from file and send them to workers
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		password := make([]byte, len(line))
		copy(password, line)
		linesCh <- password
	}
	close(linesCh)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}

	wg.Wait()

	log.Println("Finished")
}

func startProcX(outputPath string, vaults []Vault, crackedCount *int32, linesProcessed *int32, stopChan chan struct{}) {
	var writer *bufio.Writer
	if outputPath != "" {
		outputFile, err := os.OpenFile(outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Error opening output file: %v", err)
		}
		defer outputFile.Close()
		writer = bufio.NewWriter(outputFile)
	} else {
		writer = bufio.NewWriter(os.Stdout)
	}
	defer writer.Flush()

	processX(vaults, new(sync.Mutex), writer, crackedCount, linesProcessed, stopChan)

	log.Println("Finished")
}

// -w
func processPassword(password []byte, vaults []Vault, writerMu *sync.Mutex, writer *bufio.Writer,
	crackedCount *int32, linesProcessed *int32, stopChan chan struct{}) {
	atomic.AddInt32(linesProcessed, 1)

	// Try both encodings: plain (latest extension) and SHA512+base58 (older versions)
	sha512Hash := sha512.Sum512(password)
	passwordVariants := [][]byte{
		password,                             // plain password - latest Solflare extension
		[]byte(base58.Encode(sha512Hash[:])), // SHA512+base58 - older versions
	}

	for i := range vaults {
		if atomic.LoadInt32(&vaults[i].Decrypted) == 0 {
			for _, passToUse := range passwordVariants {
				decryptedData, seedPhrase, err := decryptVault(vaults[i].EncryptedData, passToUse,
					vaults[i].Salt, vaults[i].Nonce, vaults[i].Iterations, vaults[i].Kdf)
				if err != nil || !isValid(decryptedData) {
					continue
				}
				if atomic.CompareAndSwapInt32(&vaults[i].Decrypted, 0, 1) {
					output := fmt.Sprintf("Hash: %s\nPassword: %s\nSeed Phrase: %s\n",
						vaults[i].VaultText, string(password), seedPhrase)
					if writer != nil {
						writerMu.Lock()
						atomic.AddInt32(crackedCount, 1)
						writer.WriteString(output)
						writer.Flush()
						writerMu.Unlock()
					}
					if isAllVaultsCracked(vaults) {
						closeStopChannel(stopChan)
					}
					return
				}
			}
		}
	}
}

// -x
func processX(vaults []Vault, writerMu *sync.Mutex, writer *bufio.Writer,
	crackedCount *int32, linesProcessed *int32, stopChan chan struct{}) {
	atomic.AddInt32(linesProcessed, 1)

	for i := range vaults {
		if atomic.LoadInt32(&vaults[i].Decrypted) == 0 {

			if len(vaults[i].DecodedPassword) == 0 {
				continue // no vulnerability key found
			}
			passToUse := vaults[i].DecodedPassword
			decryptedData, seedPhrase, err := decryptVault(
				vaults[i].EncryptedData,
				passToUse,
				vaults[i].Salt,
				vaults[i].Nonce,
				vaults[i].Iterations,
				vaults[i].Kdf,
			)
			if err != nil || !isValid(decryptedData) {
				continue
			}
			if atomic.CompareAndSwapInt32(&vaults[i].Decrypted, 0, 1) {
				output := fmt.Sprintf("Hash: %s\nPassword: %s\nSeed Phrase: %s\n",
					vaults[i].VaultText, string("No password needed, see https://forum.hashpwn.net/post/416"), seedPhrase)
				if writer != nil {
					writerMu.Lock()
					atomic.AddInt32(crackedCount, 1)
					writer.WriteString(output)
					writer.Flush()
					writerMu.Unlock()
				}
				if isAllVaultsCracked(vaults) {
					closeStopChannel(stopChan)
				}
				return
			}
		}
	}
}
