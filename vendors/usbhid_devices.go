package vendors

import (
	"strings"

	trezorOne "github.com/xaionaro-go/cryptoWallet/internal/wallets/satoshilabs/trezor/models/1"
)

type USBHIDDevice struct {
	Device
}
type USBHIDDevices map[uint16]map[uint16]*USBHIDDevice

func GetUSBHIDDevices() USBHIDDevices {
	return USBHIDDevices{
		// SatoshiLabs
		0x534c: map[uint16]*USBHIDDevice{
			// Trezor One
			0x0001: {Device: Device{Name: "Trezor One", Factory: trezorOne.New}},
		},
	}
}

func GetVendorId(vendorName string) uint16 {
	switch strings.ToLower(vendorName) {
	case "satoshilabs", "trezor":
		return 0x534c
	}
	panic("unknown vendorName: " + vendorName)
}
