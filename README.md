[![Build Status](https://travis-ci.org/xaionaro-go/cryptoWallet.svg?branch=master)](https://travis-ci.org/xaionaro-go/cryptoWallet)
[![go report](https://goreportcard.com/badge/github.com/xaionaro-go/cryptoWallet)](https://goreportcard.com/report/github.com/xaionaro-go/cryptoWallet)
[![GoDoc](https://godoc.org/github.com/xaionaro-go/cryptoWallet?status.svg)](https://godoc.org/github.com/xaionaro-go/cryptoWallet)

Supported devices:

| Vendor | Product | Vendor ID | Product ID | Version | Notes |
| ------ | ------- | --------- | ---------- | ------- | ----- |
| SatoshiLabs | Bitcoin Wallet [TREZOR] | 0x534c | 0x0001 | 1.6.0 |Tested on Linux |

```go
[...]
wallet := cryptoWallet.FindAny()
[...]

masterKey := []byte("some key here")

encryptedMasterKey, err := wallet.EncryptKey(
	`m/3'/14'/15'/93'`,
	masterKey,
	[]byte{},
	"aWalletKeyName")
checkError(err)

decryptedMasterKey, err := wallet.DecryptKey(
	`m/3'/14'/15'/93'`,
	encryptedMasterKey,
	[]byte{},
	"aWalletKeyName")
checkError(err)

fmt.Printf("%v (%d)\n", string(masterKey), len(masterKey))
fmt.Println(encryptedMasterKey)
fmt.Printf("%v (%d)\n", string(decryptedMasterKey), len(decryptedMasterKey))
fmt.Println(decryptedMasterKey)

[...]
```
Running the example:
```
$ go run example/main.go 
some key here (13)
[167 124 140 62 203 124 234 209 28 12 1 67 101 97 228 141]
some key here (16)
[115 111 109 101 32 107 101 121 32 104 101 114 101 0 0 0]
```

A key to be encrypted/decrypted should be a multiple of 16 bytes. If you pass not a multiple of 16 bytes it will pad the value with zeros.
