package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

/*
Cyclone's Solflare Vault Decryptor
https://github.com/cyclone-github/solflare_pwn
POC tool to decrypt Solflare Vault wallets
This tool is proudly the first solflare Vault Decryptor / Cracker
coded by cyclone in Go

GNU General Public License v2.0
https://github.com/cyclone-github/solflare_pwn/blob/main/LICENSE

version history
v0.1.0; 2025-02-12;
	initial release
v0.1.1; 2025-02-20;
	tweak vulnerability logic
	clean up code
v0.2.0; 2025-03-13;
	added raw walletdata JSON input (latest extension)
	added plain + SHA512+base58 password variants (old + new)
	relaxed salt validation, flexible mnemonic extraction
v0.3.1; 2026-03-13;
	add support for latest extension
*/

// main func
func main() {
	wordlistFileFlag := flag.String("w", "", "Input file to process (omit -w to read from stdin)")
	vaultFileFlag := flag.String("h", "", "Vault File")
	xpassFlag := flag.Bool("x", false, "Exploit solflarexpass vulnerability")
	outputFile := flag.String("o", "", "Output file to write hashes to (omit -o to print to console)")
	cycloneFlag := flag.Bool("cyclone", false, "")
	versionFlag := flag.Bool("version", false, "Program version:")
	helpFlag := flag.Bool("help", false, "Prints help:")
	threadFlag := flag.Int("t", runtime.NumCPU(), "CPU threads to use (optional)")
	statsIntervalFlag := flag.Int("s", 60, "Interval in seconds for printing stats. Defaults to 60.")
	flag.Usage = func() {
		helpFunc()
	}
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

	if *vaultFileFlag == "" {
		fmt.Fprintln(os.Stderr, "-h (vault file) flags is required")
		fmt.Fprintln(os.Stderr, "Try running with -help for usage instructions")
		os.Exit(1)
	}

	if *xpassFlag && *wordlistFileFlag != "" {
		fmt.Fprintln(os.Stderr, "Error: Cannot use both -x and -w.")
		os.Exit(1)
	}

	startTime := time.Now()

	// set CPU threads
	numThreads := setNumThreads(*threadFlag)

	// variables
	var (
		crackedCount   int32
		linesProcessed int32
		wg             sync.WaitGroup
	)

	// channels
	stopChan := make(chan struct{})

	// goroutine to watch for ctrl+c
	handleGracefulShutdown(stopChan)

	// read vaults
	vaults, err := readVaultData(*vaultFileFlag)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading vault file:", err)
		os.Exit(1)
	}
	validVaultCount := len(vaults)

	// print welcome screen
	printWelcomeScreen(vaultFileFlag, wordlistFileFlag, validVaultCount, numThreads, *xpassFlag)

	// monitor status of workers
	wg.Add(1)
	go monitorPrintStats(&crackedCount, &linesProcessed, stopChan, startTime, validVaultCount, &wg, *statsIntervalFlag)

	// start the processing logic
	if *xpassFlag {
		startProcX(*outputFile, vaults, &crackedCount, &linesProcessed, stopChan)
	} else {
		startProc(*wordlistFileFlag, *outputFile, numThreads, vaults, &crackedCount, &linesProcessed, stopChan)
	}
	// close stop channel to signal all workers to stop
	closeStopChannel(stopChan)

	// wait for monitorPrintStats to finish
	wg.Wait()
}

// end code
