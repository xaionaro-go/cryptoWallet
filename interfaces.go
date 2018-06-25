package cryptoWallet

// Wallet is an abstract interface over all supported Wallets
type Wallet interface {
	// Sets a function to be called when it's required to enter a PIN or a passphrase
	SetGetPinFunc(func(title, description, ok, cancel string) ([]byte, error))

	// Sets a function to be called when it's required to get a confirmation
	SetGetConfirmFunc(func(title, description, ok, cancel string) (bool, error))

	// Resets the device and check if it is initialized. Call this
	// function before other functions to be sure that the device is in an
	// expected state.
	Reset() error

	// Checks the connection to the device and reconnects if required
	CheckConnection() error

	// Reconnect to the device
	Reconnect() error

	// Checks if the device answers correctly to a ping
	Ping() error

	// Encrypt a key using a symmetric algorithm. It should be a multiple of 16 bytes.
	EncryptKey(bip32Path string, decryptedKey []byte, nonce []byte, keyName string) ([]byte, error)

	// Decrypt a key using a symmetric algorithm. It should be a multiple of 16 bytes.
	DecryptKey(bip32Path string, encryptedKey []byte, nonce []byte, keyName string) ([]byte, error)

	// Returns a name of the device
	Name() string
}
