package cryptoWalletInterfaces

import (
	"github.com/trezor/trezord-go/usb/lowlevel"
)

// WebUSBWallet is an abstract interface over WebUSB Wallets
type WebUSBWallet interface {
	USBWallet

	// SetWebUSBDevice sets the WebUSB device to be used to reach the Trezor device
	SetWebUSBDevice(device lowlevel.Device)
}
