package cryptoWalletInterfaces

// Wallet is an abstract interface over all supported Wallets
type Wallet interface {
	// Sets a function to be called when it's required to enter a PIN or a passphrase
	SetGetPinFunc(func(title, description, ok, cancel string) ([]byte, error))

	// Sets a function to be called when it's required to get a confirmation
	SetGetConfirmFunc(func(title, description, ok, cancel string) (bool, error))

	// Call a function to get a PIN or a passphrase
	GetPin(title, description, ok, cancel string) ([]byte, error)

	// Call a function to get a confirmation
	GetConfirm(title, description, ok, cancel string) (bool, error)

	// Resets the device and check if it is initialized. Call this
	// function before other functions to be sure that the device is in an
	// expected state.
	//
	// [trezor] See also: https://doc.satoshilabs.com/trezor-tech/api-workflows.html#initialize-features
	Reset() error

	// Checks the connection to the device and reconnects if required
	CheckConnection() error

	// Reconnect tries to reconnect to find and reconnect to the device
	//
	// If the wallet is not found it calls GetConfirm method to get a confirmation
	// that it's required to try one more time.
	Reconnect() error

	// Checks if the device answers correctly to a ping
	Ping() error

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
	EncryptKey(bip32Path string, decryptedKey []byte, nonce []byte, keyName string) ([]byte, error)

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
	DecryptKey(bip32Path string, encryptedKey []byte, nonce []byte, keyName string) ([]byte, error)

	// Returns a name of the device
	Name() string
}
