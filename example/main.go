package main

import (
	"fmt"
	"github.com/xaionaro-go/cryptoWallet"
	"github.com/xaionaro-go/pinentry"
)

var (
	pinentryClient pinentry.PinentryClient
)

func getPin(title, description, ok, cancel string) ([]byte, error) {
	pinentryClient.SetPrompt(title)
	pinentryClient.SetDesc(description)
	pinentryClient.SetOK(ok)
	pinentryClient.SetCancel(cancel)
	return pinentryClient.GetPin()
}

func getConfirm(title, description, ok, cancel string) (bool, error) {
	pinentryClient.SetPrompt(title)
	pinentryClient.SetDesc(description)
	pinentryClient.SetOK(ok)
	pinentryClient.SetCancel(cancel)
	return pinentryClient.Confirm(), nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	wallet := cryptoWallet.FindAny()
	if wallet == nil {
		panic("No wallets found")
	}

	err := wallet.Ping()
	checkError(err)

	pinentryClient, err = pinentry.NewPinentryClient()
	checkError(err)

	wallet.SetGetPinFunc(getPin)
	wallet.SetGetConfirmFunc(getConfirm)

	masterKey := []byte("some key here")

	encryptedMasterKey, err := wallet.EncryptKey(`m/3'/14'/15'/93'`, masterKey, []byte{}, "someWalletKeyName")
	checkError(err)

	decryptedMasterKey, err := wallet.DecryptKey(`m/3'/14'/15'/93'`, encryptedMasterKey, []byte{}, "someWalletKeyName")
	checkError(err)

	fmt.Printf("%v (%d)\n%v\n%v (%d)\n%v\n", string(masterKey), len(masterKey), encryptedMasterKey, string(decryptedMasterKey), len(decryptedMasterKey), decryptedMasterKey)
}
