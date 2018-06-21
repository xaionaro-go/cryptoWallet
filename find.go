package cryptoWallet

type Filter struct {
	IsUSBHID   *bool
	VendorId   *uint16
	ProductIds []uint16
}

func FindAny() Wallet {
	return Find(Filter{})
}
