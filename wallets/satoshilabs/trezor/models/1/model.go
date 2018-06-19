package trezorOne

import (
	I "github.com/xaionaro-go/cryptoWallet/interfaces"
	trezorBase "github.com/xaionaro-go/cryptoWallet/wallets/satoshilabs/trezor"
	"github.com/zserge/hid"
)

type trezorOne struct {
	trezorBase.TrezorBase
}

func New(device hid.Device, name string) I.Wallet {
	instance := &trezorOne{}
	instance.TrezorBase.SetParent(instance)
	instance.SetHIDDevice(device)
	instance.SetName(name)
	return instance
}
