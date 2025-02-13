[![Readme Card](https://github-readme-stats.vercel.app/api/pin/?username=cyclone-github&repo=solflare_pwn&theme=gruvbox)](https://github.com/cyclone-github/solflare_pwn/)

[![GitHub issues](https://img.shields.io/github/issues/cyclone-github/solflare_pwn.svg)](https://github.com/cyclone-github/solflare_pwn/issues)
[![License](https://img.shields.io/github/license/cyclone-github/solflare_pwn.svg)](LICENSE)
[![GitHub release](https://img.shields.io/github/release/cyclone-github/solflare_pwn.svg)](https://github.com/cyclone-github/solflare_pwn/releases)

# Solflare Vault Extractor & Decryptor
### POC tools to recover, extract and decrypt Solflare Vaults
_**This toolset is proudly the first to announce support for Recovering Solflare Wallets**_
- `Do to many of my GitHub projects being copied and sold -- which is in violation of their GPL v2.0 License -- I will not be releasing the source code for this project (yet).`
- Contact me at https://forum.hashpwn.net/user/cyclone if you need help recovering your Solflare wallet password or seed phrase. 

### Solflare Vault location for Chrome extensions:
- Linux: `/home/$USER/.config/google-chrome/Default/Local\ Extension\ Settings/bhhhlbepdkbapadjdnnojkbgioiodbic/`
- Mac: `Library>Application Support>Google>Chrome>Default>Local Extension Settings>bhhhlbepdkbapadjdnnojkbgioiodbic`
- Windows: `C:\Users\$USER\AppData\Local\Google\Chrome\User Data\Default\Local Extension Settings\bhhhlbepdkbapadjdnnojkbgioiodbic\`
### Extractor usage example on test vault: (plaintext is "cyclone")
```
./solflare_extractor.bin bhhhlbepdkbapadjdnnojkbgioiodbic/
 ----------------------------------------------------- 
|        Cyclone's Solflare Vault Hash Extractor      |
|        Use Solflare Vault Decryptor to decrypt      |
|    https://github.com/cyclone-github/solflare_pwn   |
 ----------------------------------------------------- 

Encrypted Solflare Vault:
{"data":{"digest":"sha256","encoding":"base64","encrypted64":"iYoRET9we/aKBWtefA+6itg8izn6YMLA2WRRkctlzjmmQLAdfDT+VX9VBzuLtp/lBn0zVrs8TVR5DrgOy57+c9vuEu8ohAxdMcGWPn4FON68PEj3N616XsNe0Zcx6zMdGFpoQhatgn0H5mvtBx3vt1/HuWw6jtqX+Ftza/VWngZ62XNg4Jczq7huppxfkbyqsIvPJLoNumprZmmqfY10/Qhgbb9IrKgyYAmSCBG0QDIf3PzflOdiXVD2Baq3UAqqZ89GoB6XQdRyKp/7IG5+nnteih1IgvMTnODxVTuJHAmDm+yWqsW1dP7W+S6NkJeg1ICIUY5jcPG2ULJ6orkrJMIOqzsz8Va7UMzM87SLnhrnibht99cFuXh2V1pG7IJMds1m6ac+zoGT/SsJzooZlfCGFLdGYT30+/ozCCbSrtyNISn9Tvb9+uVPNY3m6eVWNw8FrrkRWPWw4NtyC3fw0YQpx/P8/5ctvgkM+/s63i9AnNQVWwRBQP/a4xnr9LO3vaSgH/bvhXqmnDpxwLpYkbLxprBUsTsDG6o4Ekx5zPDXKjMYBmwzgXWj1q9AJ5JBXPOnFaHLsPPNqL5zbrL9D1VFOniJoUxvqWeobAtl7GEwKb1t8DfZYjZDF4UQENqL4L7p5JVn+XCVMFLk/Vme/JnvX2MkiRLNe/cNLVlKGT7wK+0urEpa0xFUka4At1poFu2XShV3bX7JQcIR3z+aZJ2Bo7UJkgrHX03c9mgCbR7k72tG7uC62zGnDo2QKTo5A0l+aMIsmd0ZTYtDCddlwjAsIiEC4L8mH9ocf5w+qQ==","iterations":100000,"kdf":"pbkdf2","nonce64":"s7D8vko2q/za2X34xcO+S2M+r07lPTkv","salt64":"qa6AWu8JNqC0indVRgG5SA=="},"locked":true}
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
Hash: {"data":{"digest":"sha256","encoding":"base64","encrypted64":"iYoRET9we/aKBWtefA+6itg8izn6YMLA2WRRkctlzjmmQLAdfDT+VX9VBzuLtp/lBn0zVrs8TVR5DrgOy57+c9vuEu8ohAxdMcGWPn4FON68PEj3N616XsNe0Zcx6zMdGFpoQhatgn0H5mvtBx3vt1/HuWw6jtqX+Ftza/VWngZ62XNg4Jczq7huppxfkbyqsIvPJLoNumprZmmqfY10/Qhgbb9IrKgyYAmSCBG0QDIf3PzflOdiXVD2Baq3UAqqZ89GoB6XQdRyKp/7IG5+nnteih1IgvMTnODxVTuJHAmDm+yWqsW1dP7W+S6NkJeg1ICIUY5jcPG2ULJ6orkrJMIOqzsz8Va7UMzM87SLnhrnibht99cFuXh2V1pG7IJMds1m6ac+zoGT/SsJzooZlfCGFLdGYT30+/ozCCbSrtyNISn9Tvb9+uVPNY3m6eVWNw8FrrkRWPWw4NtyC3fw0YQpx/P8/5ctvgkM+/s63i9AnNQVWwRBQP/a4xnr9LO3vaSgH/bvhXqmnDpxwLpYkbLxprBUsTsDG6o4Ekx5zPDXKjMYBmwzgXWj1q9AJ5JBXPOnFaHLsPPNqL5zbrL9D1VFOniJoUxvqWeobAtl7GEwKb1t8DfZYjZDF4UQENqL4L7p5JVn+XCVMFLk/Vme/JnvX2MkiRLNe/cNLVlKGT7wK+0urEpa0xFUka4At1poFu2XShV3bX7JQcIR3z+aZJ2Bo7UJkgrHX03c9mgCbR7k72tG7uC62zGnDo2QKTo5A0l+aMIsmd0ZTYtDCddlwjAsIiEC4L8mH9ocf5w+qQ==","iterations":100000,"kdf":"pbkdf2","nonce64":"s7D8vko2q/za2X34xcO+S2M+r07lPTkv","salt64":"qa6AWu8JNqC0indVRgG5SA=="},"locked":true}
Password: cyclone
Seed Phrase: daring rose clump noble element fork differ inform gravity turtle oven iron
2025/02/12 11:14:25 Decrypted: 1/1 xxxxx.xx h/s 00h:00m:03s

2025/02/12 11:14:25 Finished
```
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
<!--
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
