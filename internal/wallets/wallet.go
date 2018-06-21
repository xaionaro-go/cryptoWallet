package internal

import (
	"github.com/zserge/hid"
)

// WalletBase is a structure to be included by implementation of real wallets.
// It implements basic routines that are used in every crypto wallet.
type WalletBase struct {
	name string

	getPin     func(title, description, ok, cancel string) ([]byte, error)
	getConfirm func(title, description, ok, cancel string) (bool, error)
}

// SetGetPinFunc sets a function to be called when it's required to enter
// a PIN or a passphrase
func (base *WalletBase) SetGetPinFunc(getPinFunc func(title, description, ok, cancel string) ([]byte, error)) {
	base.getPin = getPinFunc
}

// GetPin calls a function to get a PIN or a passphrase.
func (base *WalletBase) GetPin(title, description, ok, cancel string) ([]byte, error) {
	return base.getPin(title, description, ok, cancel)
}

// SetGetConfirmFunc sets a function to be called when it's required to get
// a confirmation. For example, this may be required to confirm "try to
// reconnect" if connection lost to the device
func (base *WalletBase) SetGetConfirmFunc(getConfirmFunc func(title, description, ok, cancel string) (bool, error)) {
	base.getConfirm = getConfirmFunc
}

// GetConfirm calls a function to get a confirmation
func (base *WalletBase) GetConfirm(title, description, ok, cancel string) (bool, error) {
	return base.getConfirm(title, description, ok, cancel)
}

// Name return a name of the device (the default value is defined in vendors/)
func (base WalletBase) Name() string {
	return base.name
}

// SetName sets a new name of the device
func (base *WalletBase) SetName(newName string) {
	base.name = newName
}

// USBHIDWalletBase is a structure to extend WalletBase for USB HID devices
type USBHIDWalletBase struct {
	WalletBase
	device      hid.Device
	vendorID    uint16
	productID   uint16
	interfaceID uint8
}

// SetHIDDevice sets an USB HID device to be used
func (base *USBHIDWalletBase) SetHIDDevice(device hid.Device) {
	base.device = device
	info := device.Info()
	base.vendorID = info.Vendor
	base.productID = info.Product
	base.interfaceID = info.Interface
}

// GetVendorID returns USB device vendor ID
func (base USBHIDWalletBase) GetVendorID() uint16 {
	return base.vendorID
}

// GetProductID returns USB device product ID
func (base USBHIDWalletBase) GetProductID() uint16 {
	return base.productID
}

// GetInterfaceID returns USB device interface ID
func (base USBHIDWalletBase) GetInterfaceID() uint8 {
	return base.interfaceID
}

// GetHIDDevice returns previously set USB HID device
func (base USBHIDWalletBase) GetHIDDevice() hid.Device {
	return base.device
}
