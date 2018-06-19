package vendors

import (
	I "github.com/xaionaro-go/cryptoWallet/interfaces"
	"github.com/zserge/hid"
)

type Device struct {
	Name    string
	Factory func(device hid.Device, name string) I.Wallet
}
