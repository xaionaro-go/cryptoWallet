package cryptoWallet

import (
	I "github.com/xaionaro-go/cryptoWallet/interfaces"
)

type Filter struct {
	IsUSBHID   *bool
	VendorId   *uint16
	ProductIds []uint16
}

func FindAny() I.Wallet {
	return Find(Filter{})
}
