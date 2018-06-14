package trezorOne

import (
	trezorBase "github.com/xaionaro-go/cryptoWallet/wallets/satoshilabs/trezor"
	"github.com/zserge/hid"
)

type trezorOne struct {
	trezorBase.TrezorBase
}

func New(device hid.Device) *trezorOne {
	instance := &trezorOne{}
	instance.TrezorBase.SetParent(instance)
	instance.SetHIDDevice(device)
	return instance
}

func (trezor trezorOne) Name() string {
	return "Trezor One"
}
