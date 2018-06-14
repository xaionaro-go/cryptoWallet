// +build !linux

package cryptoWalletRoutines

import (
	I "github.com/xaionaro-go/cryptoWallet/interfaces"
	"log"
)

func USBHIDReconnect(parent I.USBHIDWallet) error {
	log.Panic("cryptowallets are not supported on this platform :(")
}
