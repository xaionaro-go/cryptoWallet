package trezorT

import (
	trezorBase "github.com/xaionaro-go/cryptoWallet/wallets/satoshilabs/trezor"
	"github.com/zserge/hid"
)

type trezorT struct {
	trezorBase.TrezorBase
}

func New(device hid.Device) *trezorT {
	instance := &trezorT{}
	instance.TrezorBase.SetParent(instance)
	instance.SetHIDDevice(device)
	return instance
}

func (trezor trezorT) Name() string {
	return "Trezor T"
}
