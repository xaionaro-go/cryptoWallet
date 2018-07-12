// +build linux

package cryptoWallet

import (
	"github.com/conejoninja/hid"
	"github.com/trezor/trezord-go/usb/lowlevel"
	I "github.com/xaionaro-go/cryptoWallet/interfaces"
	"github.com/xaionaro-go/cryptoWallet/vendors"
)

// Find returns all known wallets that fits to the `filter`.
//
// - If `filter.VendorID` is nil then it will search for any vendor and product
//   IDs
//
// - If `filter.ProductIDs` is an empty slice and `filter.VendorID` is not nil
//   then it will search for any products of the defined vendor ID.
//
// - If the `filter` is empty then it will search for any wallets
//
// At the moment the only supported platform is Linux
func Find(filter Filter) (result []I.Wallet) {
	possibleUSBDevices := vendors.GetUSBDevices()

	// Trezor USB hid

	hid.UsbWalk(func(device hid.Device) {
		info := device.Info()
		if !filter.IsFit(info.Vendor, info.Product) {
			return
		}
		newWallet := possibleUSBDevices.GetUSBWallet(device, info.Vendor, info.Product, info.Interface)
		if newWallet == nil { // There's no known wallet with such vendorID, productID, interfaceID for this type of lower device ("hid.Device")
			return
		}
		result = append(result, newWallet.(I.Wallet))
	})

	// Trezor WebUSB

	var usbctx lowlevel.Context
	list, err := lowlevel.Get_Device_List(usbctx)
	if err != nil {
		return
	}
	for _, device := range list {
		c, err := lowlevel.Get_Active_Config_Descriptor(device)
		if err != nil {
			continue
		}

		// See https://github.com/trezor/trezord-go/blob/ecbe7156ef6bb8030cd213cd577299002dd1c409/usb/webusb.go#L217
		matches := c.BNumInterfaces > 0 &&
			c.Interface[0].Num_altsetting > 0 &&
			c.Interface[0].Altsetting[0].BInterfaceClass == lowlevel.CLASS_VENDOR_SPEC
		lowlevel.Free_Config_Descriptor(c)
		if !matches {
			continue
		}

		info, err := lowlevel.Get_Device_Descriptor(device)
		if err != nil {
			continue
		}

		newWallet := possibleUSBDevices.GetUSBWallet(device, info.IdVendor, info.IdProduct, 0)
		if newWallet == nil {
			continue
		}
		result = append(result, newWallet.(I.Wallet))
	}

	lowlevel.Free_Device_List(list, 0) // TODO: fix the memleak here

	return
}
