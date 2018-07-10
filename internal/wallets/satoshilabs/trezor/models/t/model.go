package trezorT

import (
	"github.com/conejoninja/hid"
	I "github.com/xaionaro-go/cryptoWallet/internal/interfaces"
	trezorBase "github.com/xaionaro-go/cryptoWallet/internal/wallets/satoshilabs/trezor"
)

type trezorT struct {
	trezorBase.TrezorBase
}

// New returns a new wallet "Trezor T" of vendor "SatoshiLabs"
//
// device - is a USB HID device to reach the "Trezor T"
// name - is the name from vendors/
func New(device hid.Device, name string) I.Wallet {
	instance := &trezorT{}
	instance.SetHIDDevice(device)
	instance.SetName(name)
	return instance
}
