[![Readme Card](https://github-readme-stats.vercel.app/api/pin/?username=cyclone-github&repo=solflare_pwn&theme=gruvbox)](https://github.com/cyclone-github/solflare_pwn/)

[![GitHub issues](https://img.shields.io/github/issues/cyclone-github/solflare_pwn.svg)](https://github.com/cyclone-github/solflare_pwn/issues)
[![License](https://img.shields.io/github/license/cyclone-github/solflare_pwn.svg)](LICENSE)
[![GitHub release](https://img.shields.io/github/release/cyclone-github/solflare_pwn.svg)](https://github.com/cyclone-github/solflare_pwn/releases)

# Solflare Vault Extractor & Decryptor
### POC tools to recover, extract and decrypt Solflare Vaults
_**This toolset is proudly the first to announce support for Recovering Solflare Wallets**_
- Do to many of my GitHub projects being copied and sold -- which is in violation of their GNU v2.0 License -- I will not be releasing the source code for this project.
- Contact me at https://forum.hashpwn.net/user/cyclone if you need help recovering your Solflare wallet password or seed phrase

### Solflare Vault location for Chrome extensions:
- Linux: `/home/$USER/.config/google-chrome/Default/Local\ Extension\ Settings/bhhhlbepdkbapadjdnnojkbgioiodbic/`
- Mac: `Library>Application Support>Google>Chrome>Default>Local Extension Settings>bhhhlbepdkbapadjdnnojkbgioiodbic`
- Windows: `C:\Users\$USER\AppData\Local\Google\Chrome\User Data\Default\Local Extension Settings\bhhhlbepdkbapadjdnnojkbgioiodbic\`
### Extractor usage example on test vault: (plaintext is `cyclone`)
```
./solflare_extractor.bin bhhhlbepdkbapadjdnnojkbgioiodbic/
 ----------------------------------------------------- 
|        Cyclone's Solflare Vault Hash Extractor      |
|        Use Solflare Vault Decryptor to decrypt      |
|    https://github.com/cyclone-github/solflare_pwn   |
 ----------------------------------------------------- 

Encrypted Solflare Vault:
{foobar_solflare_hash}
```
### Decryptor usage example:
```
 ----------------------------------------------- 
|       Cyclone's Solflare Vault Decryptor       |
| https://github.com/cyclone-github/solflare_pwn |
 ----------------------------------------------- 

Vault file:     hash.txt
Valid Vaults:   1
CPU Threads:    16
Wordlist:       wordlist.txt
2025/02/12 11:14:22 Working...
{foobar_solflare_hash}:cyclone
2025/02/12 11:14:25 Decrypted: 1/1 xxxxx.xx h/s 00h:00m:03s

2025/02/12 11:14:25 Finished
```
<!--
### Decryptor supported options:
```
-w {wordlist} (omit -w to read from stdin)
-h {solflare_wallet_hash}
-o {output} (omit -o to write to stdout)
-t {cpu threads}
-s {print status every nth sec}

-version (version info)
-help (usage instructions)

./solflare_decryptor.bin -h {solflare_wallet_hash} -w {wordlist} -o {output} -t {cpu threads} -s {print status every nth sec}

./solflare_decryptor.bin -h solflare.txt -w wordlist.txt -o cracked.txt -t 16 -s 10

cat wordlist | ./solflare_decryptor.bin -h solflare.txt

./solflare_decryptor.bin -h solflare.txt -w wordlist.txt -o output.txt
```
### Compile from source:
- This assumes you have Go and Git installed
  - `git clone https://github.com/cyclone-github/solflare_pwn.git`
  - solflare_extractor
  - `cd solflare_pwn/solflare_extractor`
  - `go mod init solflare_extractor`
  - `go mod tidy`
  - `go build -ldflags="-s -w" .`
  - solflare_decryptor
  - `cd solflare_pwn/solflare_decryptor`
  - `go mod init solflare_decryptor`
  - `go mod tidy`
  - `go build -ldflags="-s -w" .`
- Compile from source code how-to:
  - https://github.com/cyclone-github/scripts/blob/main/intro_to_go.txt
-->
