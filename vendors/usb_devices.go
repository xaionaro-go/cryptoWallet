package vendors

import (
	"strings"

	I "github.com/xaionaro-go/cryptoWallet/interfaces"
	trezorOne "github.com/xaionaro-go/cryptoWallet/internal/wallets/satoshilabs/trezor/models/1"
	trezorT "github.com/xaionaro-go/cryptoWallet/internal/wallets/satoshilabs/trezor/models/t"
)

// USBDevice is a factory struct for USB HID crypto wallet device
//
// Factory — an entity that creates new objects of required type
//
// Warning: Do not rely on this type, it can change in future!
type USBDevice struct {
	deviceBase
	factory           func(lowerDevice interface{}, name string) I.Wallet
	lowerDeviceSample interface{}
}

// New is a factory function for a USB HID crypto wallet device.
func (d USBDevice) New(lowerDevice interface{}) I.Wallet {
	return d.factory(lowerDevice, d.name)
}

// USBDevices just a map: map[vendorID][productID][interfaceID]*USBDevice
//
// USBDevice is a factory
//
// Warning: Do not rely on this type, it can change in future!
type USBDevices map[uint16]map[uint16]map[uint8]*USBDevice

// GetUSBWallet returns a valid cryptoWalletInterfaces.USBWallet if a
// known combination of (the lowerDevice type, the vendorID value, the
// productID value and the interfaceID value) is passed
//
// If the combination is unknown then nil is returned.
//
// Note about the lowerDevice type: Trezor One will be returned only if
// lowerDevice is of type hid.Device; Trezor T will be returned only if
// lowerDevice is of type lowlevel.Device
func (usbDevices USBDevices) GetUSBWallet(lowerDevice interface{}, vendorID, productID uint16, interfaceID uint8) I.USBWallet {
	if usbDevices[vendorID] == nil {
		return nil
	}
	if usbDevices[vendorID][productID] == nil {
		return nil
	}
	if usbDevices[vendorID][productID][interfaceID] == nil {
		return nil
	}
	usbDevice := usbDevices[vendorID][productID][interfaceID]
	return usbDevice.New(lowerDevice).(I.USBWallet)
}

// GetUSBDevices returns a map of a vendorID, a productID and interfaceID
// to a object factory to a crypto wallet device
//
// example:
//
// ```
//
// deviceMeta := GetUSBDevices()[0x534c][0x0001][0x00]
//
// deviceMeta.Factory(hidDevice, deviceMeta.Name)
//
// ```
//
// Warning: Do not rely on type USBDevices, it can change in future!
func GetUSBDevices() USBDevices {
	return USBDevices{
		// SatoshiLabs
		0x534c: map[uint16]map[uint8]*USBDevice{
			// Trezor One
			0x0001: {
				// See https://github.com/trezor/trezor-mcu/blob/826b764085c0e637eed5bf631bacea964327289d/firmware/usb.c#L32
				0x00: {deviceBase: deviceBase{name: "Trezor One"}, factory: trezorOne.New}, // the main interface (USB HID)
				//0x01: , // U2F interface
			},
		},
		0x1209: map[uint16]map[uint8]*USBDevice{
			// Trezor T
			0x53C1: {
				0x00: {deviceBase: deviceBase{name: "Trezor T"}, factory: trezorT.New}, // the main interface (WebUSB)
				//0x01: , // USB HID (special purposes)
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
