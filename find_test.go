package cryptoWallet

import (
	"github.com/xaionaro-go/cryptoWallet/vendors"
)

const (
	dummyKey          = "example-example-"
	nonce             = "" // the "nonce" is optional, so it's empty for now
	keyName           = "cryptoWalletExample"
	keyDerivationPath = `m/11019'/0'/0'/0'`
)

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
	// Some handler here that will return valid PIN-code
	return []byte{}, nil
}
func getConfirm(title, description, ok, cancel string) (bool, error) {
	// Some handler here that will return true or false
	return false, nil
}
func doSomething(...interface{}) {}

// An example for FindAny()
func ExampleFindAny() {
	// Getting a wallet
	wallet := FindAny()
	if wallet == nil {
		// The device wasn't found
		// Process the error here
		return
	}

	// Put the device to the initial state. This also checks if the device
	// is initialized
	err := wallet.Reset()
	if err != nil {
		// Process the error here. For example, if device is not
		// initialized (ErrNotInitialized)
		return
	}

	// Setting required handlers
	wallet.SetGetPinFunc(getPin)
	wallet.SetGetConfirmFunc(getConfirm)

	// Setting a key to check
	masterKey := []byte("some key here")

	// Encrypting it. The encrypted value will be of length 16 bytes (instead of 13 bytes).
	encryptedMasterKey, err := wallet.EncryptKey(`m/3'/14'/15'`, masterKey, []byte{}, "aWalletKeyName")
	if err != nil {
		// Cannot encrypt the key. For example, it may happen if wrong PIN code was passed to the device
		// Process error here
		return
	}

	// Decrypting it back. Now it will be of length 16 bytes (instead of 13 bytes).
	decryptedMasterKey, err := wallet.DecryptKey(`m/3'/14'/15'`, encryptedMasterKey, []byte{}, "aWalletKeyName")
	if err != nil {
		// Cannot decrypt the key.
		// Process error here
		return
	}

	// At the moment: string(masterKey) == string(decryptedMasterKey[:len(masterKey)])
	if string(decryptedMasterKey[:len(masterKey)]) != string(masterKey) {
		// Something is wrong
	}
}

// An example for Find()
//
// This example is adapted from gocryptfs
func ExampleFind() {
	// This example shows how to generate a deterministic key using "Trezor One"

	// Find all trezor devices
	trezors := Find(Filter{
		VendorID:   &[]uint16{vendors.GetVendorID("satoshilabs")}[0],
		ProductIDs: []uint16{1 /* Trezor One */},
	})

	// For example we require to one and only one trezor device to be connected.
	if len(trezors) == 0 {
		// Trezor device is not found.
		// Process the error here
		return
	}
	if len(trezors) > 1 {
		// It's more than one Trezor device connected.
		// Process the error here
		return
	}

	// Using the first found device
	trezor := trezors[0] // also you can cast it into `cryptoWalletInterfaces.Trezor`

	// Put the device to the initial state. This also checks if the device
	// is initialized
	err := trezor.Reset()
	if err != nil {
		// Process the error here. For example, if device is not
		// initialized (ErrNotInitialized)
		return
	}

	// Trezor may ask for PIN or Passphrase. Setting the handler for this case.
	trezor.SetGetPinFunc(getPin)

	// In some cases (like lost connection to the Trezor device and cannot
	// reconnect) it's required to get a confirmation from the user to
	// retry to reconnect. Setting the handler for this case.
	trezor.SetGetConfirmFunc(getConfirm)

	// To generate a deterministic key we trying to decrypt our
	// predefined constant key using the Trezor device. The resulting key
	// will depend on next variables:
	// * the Trezor master key;
	// * the passphrase (passed to the Trezor).
	//
	// The right key will be received only if both values (mentioned
	// above) are correct.
	//
	// Note:
	// Also the resulting key depends on this values (that we defined as
	// constants above):
	// * the key derivation path;
	// * the "encrypted" key;
	// * the nonce;
	// * the key name.
	key, err := trezor.DecryptKey(keyDerivationPath, []byte(dummyKey), []byte(nonce), keyName)
	if err != nil {
		// Cannot decrypt the key. For example, it may happen if wrong PIN code was passed to the device
		// Process the error here
		return
	}

	// Everything ok: key - is the wanted key
	doSomething(key)
}
