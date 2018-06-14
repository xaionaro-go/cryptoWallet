// +build !linux

package cryptoWallet

import (
	I "github.com/xaionaro-go/cryptoWallet/interfaces"
	"log"
)

func FindAny() (result I.Wallet) {
	log.Panic("cryptowallets are not supported on this platform :(")
	return nil
}
