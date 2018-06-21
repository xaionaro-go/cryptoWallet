package cryptoWallet

/* // Using a pinentry:

import "github.com/xaionaro-go/pinentry"
var pinentryClient pinentry.PinentryClient

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
}*/

// Just a placeholders:
func getPin(title, description, ok, cancel string) ([]byte, error) {
	return []byte{}, nil
}
func getConfirm(title, description, ok, cancel string) (bool, error) {
	return false, nil
}

// To do not write "if" everytime
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func ExampleFindAny() {
	// Getting a wallet
	wallet := FindAny()
	if wallet == nil {
		return
	}

	// Setting required handlers
	wallet.SetGetPinFunc(getPin)
	wallet.SetGetConfirmFunc(getConfirm)

	// Setting a key to check
	masterKey := []byte("some key here")

	// Encrypting it. The encrypted value will be of length 16 bytes (instead of 13 bytes).
	encryptedMasterKey, err := wallet.EncryptKey(`m/3'/14'/15'`, masterKey, []byte{}, "aWalletKeyName")
	checkError(err)

	// Decrypting it back. Now it will be of length 16 bytes (instead of 13 bytes).
	decryptedMasterKey, err := wallet.DecryptKey(`m/3'/14'/15'`, encryptedMasterKey, []byte{}, "aWalletKeyName")
	checkError(err)

	// ATM, if PIN was correct then: string(masterKey) == string(decryptedMasterKey[:len(masterKey)])
	if string(decryptedMasterKey[:len(masterKey)]) != string(masterKey) {
		// Something is wrong
	}
}
