package cryptoWallet

import (
	"testing"

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

func ExampleFindAny() {
	wallet := FindAny()
	if wallet == nil {
		return
	}

	err := wallet.Ping()
	checkError(err)

	pinentryClient, err = pinentry.NewPinentryClient()
	checkError(err)

	wallet.SetGetPinFunc(getPin)
	wallet.SetGetConfirmFunc(getConfirm)

	masterKey := []byte("some key here")

	encryptedMasterKey, err := wallet.EncryptKey(`m/3'/14'/15'`, masterKey, []byte{}, "aWalletKeyName")
	checkError(err)

	decryptedMasterKey, err := wallet.DecryptKey(`m/3'/14'/15'`, encryptedMasterKey, []byte{}, "aWalletKeyName")
	checkError(err)

	if string(masterKey) != string(decryptedMasterKey[:len(masterKey)]) {
		return
	}

	return
}

func TestFindAny(t *testing.T) {
	ExampleFindAny()
}
