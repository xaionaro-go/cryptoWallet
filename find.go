package cryptoWallet

// Filter is a struct that could be passed to Find() to select devices
type Filter struct {
	IsUSBHID   *bool
	VendorID   *uint16
	ProductIDs []uint16
}

// FindAny returns any found known wallet
func FindAny() Wallet {
	result := Find(Filter{})
	if len(result) == 0 {
		return nil
	}
	return result[0]
}
