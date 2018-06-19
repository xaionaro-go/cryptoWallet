// +build !linux

package cryptoWallet

import (
	I "github.com/xaionaro-go/cryptoWallet/interfaces"
	"log"
)

func Find(filter Filter) I.Wallet {
	log.Panic("cryptowallets are not supported on this platform :(")
	return nil
}
