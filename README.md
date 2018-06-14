see `example/`:
```
...

masterKey := []byte("some key here")

encryptedMasterKey, err := wallet.EncryptKey(`m/3'/14'/15'/93'`, masterKey, []byte{}, "someWalletKeyName")
checkError(err)

decryptedMasterKey, err := wallet.DecryptKey(`m/3'/14'/15'/93'`, encryptedMasterKey, []byte{}, "someWalletKeyName")
checkError(err)

fmt.Printf("%v (%d)\n%v\n%v (%d)\n%v\n", string(masterKey), len(masterKey), encryptedMasterKey, string(decryptedMasterKey), len(decryptedMasterKey), decryptedMasterKey)
...
```

```
$ go run example/main.go 
some key here (13)
[167 124 140 62 203 124 234 209 28 12 1 67 101 97 228 141]
some key here (16)
[115 111 109 101 32 107 101 121 32 104 101 114 101 0 0 0]
```
