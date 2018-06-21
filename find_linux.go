// +build linux

package cryptoWallet

import (
	"github.com/xaionaro-go/cryptoWallet/vendors"
	"github.com/zserge/hid"
)

func Find(filter Filter) (result Wallet) {
	if filter.IsUSBHID != nil {
		if *filter.IsUSBHID != true {
			return
		}
	}
	possibleUSBHIDDevices := vendors.GetUSBHIDDevices()
	wantedProductId := map[uint16]bool{}
	for _, productId := range filter.ProductIds {
		wantedProductId[productId] = true
	}

	hid.UsbWalk(func(device hid.Device) {
		if result != nil {
			return
		}
		info := device.Info()
		if filter.VendorId != nil {
			if info.Vendor != *filter.VendorId {
				return
			}
		}
		if len(filter.ProductIds) > 0 {
			if !wantedProductId[info.Product] {
				return
			}
		}
		if possibleUSBHIDDevices[info.Vendor] == nil {
			return
		}
		if possibleUSBHIDDevices[info.Vendor][info.Product] == nil {
			return
		}
		deviceMeta := possibleUSBHIDDevices[info.Vendor][info.Product]
		result = deviceMeta.Factory(device, deviceMeta.Name)
	})

	return
}
