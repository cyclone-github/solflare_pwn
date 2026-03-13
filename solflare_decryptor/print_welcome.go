package main

import (
	"fmt"
	"log"
	"os"
)

// version func
func versionFunc() {
	fmt.Fprintln(os.Stderr, "Cyclone's Solflare Vault Decryptor v0.3.1; 2026-03-13\nhttps://github.com/cyclone-github/solflare_pwn")
	fmt.Fprintln(os.Stderr)
}

// help func
func helpFunc() {
	versionFunc()
	str := `Example Usage:

-w {wordlist} (omit -w to read from stdin)
-h {vault_file} - supports $solflare$ hash format OR raw walletdata JSON
-o {output} (omit -o to write to stdout)
-t {cpu threads}
-s {print status every nth sec}
-x {xpass exploit mode}

Vault file formats:
  - $solflare$ hash (extractor output)
  - Raw JSON: walletdata: {"data":{"encrypted64":"...","nonce64":"...","salt64":"...","iterations":...}}

./solflare_decryptor.bin -h solflare.txt -w wordlist.txt -o cracked.txt -t 16 -s 10

cat wordlist | ./solflare_decryptor.bin -h solflare.txt`
	fmt.Fprintln(os.Stderr, str)
}

// print welcome screen
func printWelcomeScreen(vaultFileFlag, wordlistFileFlag *string, validVaultCount, numThreads int, xpass bool) {
	fmt.Fprintln(os.Stderr, " ----------------------------------------------- ")
	fmt.Fprintln(os.Stderr, "|       Cyclone's Solflare Vault Decryptor       |")
	fmt.Fprintln(os.Stderr, "| https://github.com/cyclone-github/solflare_pwn |")
	fmt.Fprintln(os.Stderr, " ----------------------------------------------- ")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintf(os.Stderr, "Vault file:\t%s\n", *vaultFileFlag)
	fmt.Fprintf(os.Stderr, "Valid Vaults:\t%d\n", validVaultCount)
	fmt.Fprintf(os.Stderr, "CPU Threads:\t%d\n", numThreads)

	if *wordlistFileFlag == "" {
		if xpass {
			fmt.Fprintln(os.Stderr, "Mode:\t\txpass exploit")
		} else {
			fmt.Fprintln(os.Stderr, "Wordlist:\tReading stdin")
		}
	} else {
		fmt.Fprintf(os.Stderr, "Wordlist:\t%s\n", *wordlistFileFlag)
	}

	log.Println("Working...")
}
