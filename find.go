package cryptoWallet

import (
	I "github.com/xaionaro-go/cryptoWallet/interfaces"
)

// Filter is a struct that could be passed to Find() to select devices
type Filter struct {
	IsUSBHID   *bool
	VendorID   *uint16
	ProductIDs []uint16
}

// FindAny returns any found known wallet
func FindAny() I.Wallet {
	result := Find(Filter{})
	if len(result) == 0 {
		return nil
	}
	return result[0]
}
