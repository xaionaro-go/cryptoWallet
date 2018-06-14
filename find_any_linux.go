// +build linux

package cryptoWallet

import (
	I "github.com/xaionaro-go/cryptoWallet/interfaces"
	trezorOne "github.com/xaionaro-go/cryptoWallet/wallets/satoshilabs/trezor/models/1"
	"github.com/zserge/hid"
)

func FindAny() (result I.Wallet) {
	hid.UsbWalk(func(device hid.Device) {
		if result != nil {
			return
		}
		info := device.Info()
		switch info.Vendor {
		case 21324: // SatoshiLabs
			switch info.Product {
			case 1: // Trezor 1
				result = trezorOne.New(device)
			}
		}
	})
	return
}
