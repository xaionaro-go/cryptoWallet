package internal

import (
	"github.com/zserge/hid"
)

type WalletBase struct {
	getPin     func(title, description, ok, cancel string) ([]byte, error)
	getConfirm func(title, description, ok, cancel string) (bool, error)
}

func (base *WalletBase) SetGetPinFunc(getPinFunc func(title, description, ok, cancel string) ([]byte, error)) {
	base.getPin = getPinFunc
}

func (base *WalletBase) GetPin(title, description, ok, cancel string) ([]byte, error) {
	return base.getPin(title, description, ok, cancel)
}

func (base *WalletBase) SetGetConfirmFunc(getConfirmFunc func(title, description, ok, cancel string) (bool, error)) {
	base.getConfirm = getConfirmFunc
}

func (base *WalletBase) GetConfirm(title, description, ok, cancel string) (bool, error) {
	return base.getConfirm(title, description, ok, cancel)
}

type USBHIDWalletBase struct {
	WalletBase
	device      hid.Device
	vendorId    uint16
	productId   uint16
	interfaceId uint8
}

func (base *USBHIDWalletBase) SetDevice(device hid.Device) {
	base.device = device
	info := device.Info()
	base.vendorId = info.Vendor
	base.productId = info.Product
	base.interfaceId = info.Interface
}

func (base USBHIDWalletBase) GetVendorId() uint16 {
	return base.vendorId
}

func (base USBHIDWalletBase) GetProductId() uint16 {
	return base.productId
}

func (base USBHIDWalletBase) GetInterfaceId() uint8 {
	return base.interfaceId
}

func (base USBHIDWalletBase) GetDevice() hid.Device {
	return base.device
}
