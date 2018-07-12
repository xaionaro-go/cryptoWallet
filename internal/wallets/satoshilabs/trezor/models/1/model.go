// +build linux,cgo darwin,!ios,cgo windows,cgo

// The building requirements above a copied from
// github.com/trezor/usbhid/libusb.go (commit: 519ec1000beb862bbe9b16b99d782ab77787ea18)

package trezorOne

import (
	"fmt"

	"github.com/conejoninja/hid"
	"github.com/conejoninja/tesoro/pb/messages"
	"github.com/conejoninja/tesoro/transport"
	I "github.com/xaionaro-go/cryptoWallet/interfaces"
	routines "github.com/xaionaro-go/cryptoWallet/internal/routines"
	trezorBase "github.com/xaionaro-go/cryptoWallet/internal/wallets/satoshilabs/trezor"
)

type trezorOne struct {
	trezorBase.TrezorBase
	device hid.Device
}

// New returns a new wallet "Trezor One" of vendor "SatoshiLabs"
//
// device - is a USB HID device to reach the "Trezor One"
// name - is the name from vendors/
func New(device interface{}, name string) I.Wallet {
	instance := &trezorOne{}
	instance.SetUSBHIDDevice(device.(hid.Device))
	instance.SetName(name)
	instance.SetWallet(instance)
	return instance
}

// SetUSBHIDDevice sets USB HID device to be used to reach the crypto wallet
func (trezor *trezorOne) SetUSBHIDDevice(device hid.Device) {
	var t transport.TransportHID
	t.SetDevice(device)
	trezor.Client.SetTransport(&t)
	info := device.Info()
	trezor.SetUSBInfo(info.Vendor, info.Product, info.Interface)
	trezor.device = device
}

// getUSBHIDDevice returns previously set USB HID device
func (trezor trezorOne) getUSBHIDDevice() hid.Device {
	return trezor.device
}

// Ping checks if the device answers correctly to a ping
func (trezor *trezorOne) Ping() error {
	if trezor.getUSBHIDDevice() == nil {
		return fmt.Errorf("trezor.getUSBHIDDevice() == nil")
	}
	if _, err := trezor.getUSBHIDDevice().HIDReport(); err != nil {
		return err
	}
	pongMsg, msgType := trezor.TrezorBase.Ping("ping")
	if pongMsg == "ping" {
		return nil
	}
	switch msgType {
	case messages.MessageType_MessageType_Success:
		return fmt.Errorf("The wallet device seems to be not initialized")
	}
	return fmt.Errorf("An unexpected behaviour of the wallet device: %v: %v", msgType, pongMsg)
}

// Reconnect tries to reconnect to find and reconnect to the  USB HID device
// of the wallet.
//
// If the wallet is not found it calls GetConfirm method to get a confirmation
// that it's required to try one more time.
func (trezor *trezorOne) Reconnect() error {
	return routines.USBHIDReconnect(trezor)
}
