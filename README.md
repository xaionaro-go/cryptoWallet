Supported devices:
| Vendor | Product | Vendor ID | Product ID | Version | Notes |
| ------ | ------- | --------- | ---------- | ------- | ----- |
| SatoshiLabs | Bitcoin Wallet [TREZOR] | 0x534c | 0x0001 | 1.6.0 |Tested on Linux |

See `example/`:
```go
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
