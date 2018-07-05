// +build !linux

package cryptoWallet

import (
	"log"

	"github.com/xaionaro-go/cryptoWallet/internal/errors"
)

func Find(filter Filter) []Wallet {
	log.Panic(errors.ErrNotSupportedPlatform)
	return nil
}
