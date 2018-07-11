package cryptoWalletInterfaces

import (
	"github.com/conejoninja/hid"
)

// USBHIDWallet is an abstract interface over USB HID Wallets
type USBHIDWallet interface {
	Wallet

	// SetHIDDevice sets the USBHID device to be used to reach the Trezor device
	SetHIDDevice(device hid.Device)

	// GetProductID returns USB device product ID
	GetProductID() uint16

	// GetVendorID returns USB device vendor ID
	GetVendorID() uint16

	// GetInterfaceID returns USB device interface ID
	GetInterfaceID() uint8
}
