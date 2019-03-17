package cryptoWallet

import (
	I "github.com/xaionaro-go/cryptoWallet/interfaces"
	"github.com/xaionaro-go/cryptoWallet/internal/wallets/dummy"
)

// NewDummy returns a dummy wallet (that only imitates if he do anything; don't
// use as a real security tool, only for debugging).
func NewDummy() I.Wallet {
	return dummy.New()
}
