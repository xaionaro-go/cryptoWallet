package satoshilabsWallet

import (
	base "github.com/xaionaro-go/cryptoWallet/internal/wallets"
)

// Base is a structure to be included by implementation of wallets submitted
// by a vendor "SatoshiLabs".
type Base struct {
	base.USBWalletBase
}
