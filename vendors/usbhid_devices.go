package vendors

import (
	"strings"

	I "github.com/xaionaro-go/cryptoWallet/internal/interfaces"
	trezorOne "github.com/xaionaro-go/cryptoWallet/internal/wallets/satoshilabs/trezor/models/1"
	"github.com/conejoninja/hid"
)

// USBHIDDevice is a factory struct for USB HID crypto wallet device
//
// Factory — an entity that creates new objects of required type
//
// Warning: Do not rely on this type, it can change in future!
type USBHIDDevice struct {
	deviceBase
	factory func(device hid.Device, name string) I.Wallet
}

// New is a factory function for a USB HID crypto wallet device.
func (d USBHIDDevice) New(device hid.Device) I.Wallet {
	return d.factory(device, d.name)
}

// USBHIDDevices just a map: map[vendorID][productID][interfaceID]*USBHIDDevice
//
// USBHIDDevice is a factory
//
// Warning: Do not rely on this type, it can change in future!
type USBHIDDevices map[uint16]map[uint16]map[uint8]*USBHIDDevice

// GetUSBHIDDevices returns a map of a vendorID, a productID and interfaceID
// to a object factory to a crypto wallet device
//
// example:
//
// ```
//
// deviceMeta := GetUSBHIDDevices()[0x534c][0x0001][0x00]
//
// deviceMeta.Factory(hidDevice, deviceMeta.Name)
//
// ```
//
// Warning: Do not rely on type USBHIDDevices, it can change in future!
func GetUSBHIDDevices() USBHIDDevices {
	return USBHIDDevices{
		// SatoshiLabs
		0x534c: map[uint16]map[uint8]*USBHIDDevice{
			// Trezor One
			0x0001: {
				// See https://github.com/trezor/trezor-mcu/blob/826b764085c0e637eed5bf631bacea964327289d/firmware/usb.c#L32
				0x00: {deviceBase: deviceBase{name: "Trezor One"}, factory: trezorOne.New}, // the main interface
				//0x01: , // U2F interface
			},
		},
	}
}

// GetVendorID returns a vendorID of an USB device using vendor name for known
// devices
func GetVendorID(vendorName string) uint16 {
	switch strings.ToLower(vendorName) {
	case "satoshilabs", "trezor":
		return 0x534c
	}
	panic("unknown vendorName: " + vendorName)
}
