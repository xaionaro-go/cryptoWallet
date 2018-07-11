// +build linux

package cryptoWallet

import (
	"github.com/conejoninja/hid"
	I "github.com/xaionaro-go/cryptoWallet/interfaces"
	"github.com/xaionaro-go/cryptoWallet/vendors"
)

// Find returns all known wallets that fits to the `filter`.
//
// - If `filter.IsUSBHID` is nil then it will search for both USB HID and not
// USB HID devices
//
// - If `filter.VendorID` is nil then it will search for any vendor and product
// IDs
//
// - If `filter.ProductIDs` is an empty slice and `filter.VendorID` is not nil
// then it will search for any products of the defined vendor ID.
//
// - If the `filter` is empty then it will search for any wallets
//
// At the moment the only supported platform is Linux
func Find(filter Filter) (result []I.Wallet) {
	if filter.IsUSBHID != nil {
		if *filter.IsUSBHID != true {
			return
		}
	}
	possibleUSBHIDDevices := vendors.GetUSBHIDDevices()
	wantedProductID := map[uint16]bool{}
	for _, productID := range filter.ProductIDs {
		wantedProductID[productID] = true
	}

	hid.UsbWalk(func(device hid.Device) {
		info := device.Info()
		if filter.VendorID != nil {
			if info.Vendor != *filter.VendorID {
				return
			}
			if len(filter.ProductIDs) > 0 {
				if !wantedProductID[info.Product] {
					return
				}
			}
		}
		if possibleUSBHIDDevices[info.Vendor] == nil {
			return
		}
		if possibleUSBHIDDevices[info.Vendor][info.Product] == nil {
			return
		}
		if possibleUSBHIDDevices[info.Vendor][info.Product][info.Interface] == nil {
			return
		}
		result = append(result, possibleUSBHIDDevices[info.Vendor][info.Product][info.Interface].New(device))
	})

	return
}
