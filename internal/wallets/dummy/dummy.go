package dummy

import (
	I "github.com/xaionaro-go/cryptoWallet/interfaces"
	"github.com/xaionaro-go/cryptoWallet/internal/wallets"
)

const (
	dummyKeyLength = 16
)

var (
	dummyMasterKey = []byte(`testtesttesttest`)
)

var _ I.Wallet = &dummy{} // For compile-time check of the interface complience

// New returns a new dummy wallet
func New() I.Wallet {
	wallet := &dummy{}
	copy(wallet.masterKey[:], dummyMasterKey)
	return wallet
}

type dummy struct {
	internal.WalletBase
	masterKey [dummyKeyLength]byte
}

// Checks the connection to the device and reconnects if required
//
// Always returns nil (because it's a dummy device)
func (w *dummy) CheckConnection() error {
	return nil
}

// DecryptKey decrypts a key using a symmetric algorithm. The key length
// should be a multiple of 16 bytes.
//
// - `path` is a BIP32 path;
//
// - `encryptedKey` is a key to be decrypted;
//
// - `nonce` is "number that can only be used once",
//    see https://en.wikipedia.org/wiki/Cryptographic_nonce;
//
// - `trezorKeyname` is a key name that affects on encrypts and displays
//    on the screen of a trezor device.
func (w *dummy) DecryptKey(path string, encryptedKey []byte, nonce []byte, trezorKeyname string) ([]byte, error) {
	w.GetPin("Passphrase", "", "Confirm", "Cancel")

	var key [dummyKeyLength]byte

	copy(key[:], w.masterKey[:])
	for idx, c := range path {
		key[idx%dummyKeyLength] ^= byte(c)
	}

	for idx, c := range encryptedKey {
		key[idx%dummyKeyLength] ^= c
	}

	for idx, c := range nonce {
		key[idx%dummyKeyLength] ^= c
	}

	for idx, c := range trezorKeyname {
		key[idx%dummyKeyLength] ^= byte(c)
	}

	return key[:], nil
}

// EncryptKey encrypts a key using a symmetric algorithm. The key length
// should be a multiple of 16 bytes.
//
// - `path` is a BIP32 path;
//
// - `decryptedKey` is a key to be encrypted;
//
// - `nonce` is optional "number that can only be used once",
//    see https://en.wikipedia.org/wiki/Cryptographic_nonce;
//
// - `trezorKeyname` is a key name that affects on encrypts and displays
//    on the screen of a trezor device.
func (w *dummy) EncryptKey(path string, decryptedKey []byte, nonce []byte, trezorKeyname string) ([]byte, error) {
	return w.DecryptKey(path, decryptedKey, nonce, trezorKeyname)
}

// Checks if the device answers correctly to a ping
//
// Always returns nil (because it's a dummy device)
func (w *dummy) Ping() error {
	return nil
}

// Reconnect tries to reconnect to find and reconnect to the device
//
// Always returns nil (because it's a dummy device)
func (w *dummy) Reconnect() error {
	return nil
}

// Resets the device and check if it is initialized. Call this
// function before other functions to be sure that the device is in an
// expected state.
//
// Always returns nil (because it's a dummy device)
func (w *dummy) Reset() error {
	return nil
}
