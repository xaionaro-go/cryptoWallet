package trezorOne

import (
	trezorBase "github.com/xaionaro-go/cryptoWallet/wallets/satoshilabs/trezor"
	I "github.com/xaionaro-go/cryptoWallet/interfaces"
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
