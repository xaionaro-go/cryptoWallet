// +build !linux

package cryptoWallet

import (
	"log"
)

func Find(filter Filter) []Wallet {
	log.Panic("cryptowallets are not supported on this platform :(")
	return nil
}
