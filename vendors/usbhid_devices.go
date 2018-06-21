package vendors

import (
	"strings"

	I "github.com/xaionaro-go/cryptoWallet/internal/interfaces"
	trezorOne "github.com/xaionaro-go/cryptoWallet/internal/wallets/satoshilabs/trezor/models/1"
	"github.com/zserge/hid"
)

// USBHIDDevice is a factory struct for USB HID crypto wallet device
//
// Factory — an entity that creates new objects of required type
type USBHIDDevice struct {
	deviceBase
	factory func(device hid.Device, name string) I.Wallet
}

// New is a factory function for a USB HID crypto wallet device.
func (d USBHIDDevice) New(device hid.Device) I.Wallet {
	return d.factory(device, d.name)
}

// USBHIDDevices just a map: map[vendorID][productID]*USBHIDDevice
// USBHIDDevice is a factory
type USBHIDDevices map[uint16]map[uint16]*USBHIDDevice

// GetUSBHIDDevices returnes a map of a vendorID and a productID to a object
// factory to a crypto wallet device
//
// example:
// ```
// deviceMeta := GetUSBHIDDevices()[0x534c][0x0001]
// deviceMeta.Factory(hidDevice, deviceMeta.Name)
// ```
func GetUSBHIDDevices() USBHIDDevices {
	return USBHIDDevices{
		// SatoshiLabs
		0x534c: map[uint16]*USBHIDDevice{
			// Trezor One
			0x0001: {deviceBase: deviceBase{name: "Trezor One"}, factory: trezorOne.New},
		},
	}
}

// GetVendorID returns vendorID of an USB device using vendor name for known
// devices
func GetVendorID(vendorName string) uint16 {
	switch strings.ToLower(vendorName) {
	case "satoshilabs", "trezor":
		return 0x534c
	}
	panic("unknown vendorName: " + vendorName)
}
