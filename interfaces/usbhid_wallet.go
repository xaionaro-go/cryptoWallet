package cryptoWalletInterfaces

import (
	"github.com/conejoninja/hid"
)

// USBHIDWallet is an abstract interface over USB HID Wallets
type USBHIDWallet interface {
	USBWallet

	// SetUSBHIDDevice sets the USBHID device to be used to reach the Trezor device
	SetUSBHIDDevice(device hid.Device)
}
