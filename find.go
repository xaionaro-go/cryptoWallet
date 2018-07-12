package cryptoWallet

import (
	I "github.com/xaionaro-go/cryptoWallet/interfaces"
)

// Filter is a struct that could be passed to Find() to select devices
type Filter struct {
	VendorID   *uint16
	ProductIDs []uint16
}

// IsFit returns true if vendorID and productID matches the rules of
// the Filter struct
//
// - If `filter.VendorID` is nil then it will match to any vendorID
//   and any productID
//
// - If `filter.ProductIDs` is an empty slice and `filter.VendorID` is
//   not nil then it will match for any product of the defined vendor
//   ID.
//
// - If the `filter` is empty then it will match for vendor ID and any
//   productID
func (filter Filter) IsFit(vendorID uint16, productID uint16) bool {
	if filter.VendorID == nil {
		return true
	}

	if vendorID != *filter.VendorID {
		return false
	}

	if len(filter.ProductIDs) == 0 {
		return true
	}

	for _, wantedProductID := range filter.ProductIDs { // TODO: don't iterate everytime through this slice
		if productID == wantedProductID {
			return true
		}
	}
	return false
}

// FindAny returns any found known wallet
func FindAny() I.Wallet {
	result := Find(Filter{})
	if len(result) == 0 {
		return nil
	}
	return result[0]
}
