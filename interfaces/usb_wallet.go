package cryptoWalletInterfaces

// USBWallet is an abstract interface over USB Wallets
type USBWallet interface {
	Wallet

	// GetProductID returns USB device product ID
	GetProductID() uint16

	// GetVendorID returns USB device vendor ID
	GetVendorID() uint16

	// GetInterfaceID returns USB device interface ID
	//
	// Not supposed to be used on WebUSB devices. It always returns 0
	// in the case.
	GetInterfaceID() uint8
}
