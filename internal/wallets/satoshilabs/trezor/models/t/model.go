package trezorT

import (
	"fmt"

	"github.com/conejoninja/tesoro/pb/messages"
	"github.com/conejoninja/tesoro/transport"
	"github.com/trezor/trezord-go/usb/lowlevel"
	I "github.com/xaionaro-go/cryptoWallet/interfaces"
	routines "github.com/xaionaro-go/cryptoWallet/internal/routines"
	trezorBase "github.com/xaionaro-go/cryptoWallet/internal/wallets/satoshilabs/trezor"
)

type trezorT struct {
	trezorBase.TrezorBase
	device lowlevel.Device
}

// New returns a new wallet "Trezor T" of vendor "SatoshiLabs"
//
// device - is a WebUSB device to reach the "Trezor T"
// name - is the name from vendors/
//
// if "device" is of incorrect type then returns nil
func New(deviceI interface{}, name string) I.Wallet {
	device, ok := deviceI.(lowlevel.Device)
	if !ok {
		return nil
	}
	instance := &trezorT{}
	instance.SetWebUSBDevice(device)
	instance.SetName(name)
	return instance
}

// SetWebUSBDevice sets WebUSB device to be used to reach the crypto wallet
func (trezor *trezorT) SetWebUSBDevice(device lowlevel.Device) {
	var t transport.TransportWebUSB
	t.SetDevice(device)
	trezor.Client.SetTransport(&t)
	info, err := lowlevel.Get_Device_Descriptor(device)
	if err != nil {
		panic(err)
	}
	trezor.SetUSBInfo(info.IdVendor, info.IdProduct, 0)
	trezor.device = device
}

// getWebUSBDevice returns previously set USB  device
func (trezor trezorT) getWebUSBDevice() lowlevel.Device {
	return trezor.device
}

// Ping checks if the device answers correctly to a ping
func (trezor *trezorT) Ping() error {
	if trezor.getWebUSBDevice() == nil {
		return fmt.Errorf("trezor.getDevice() == nil")
	}
	device := trezor.getWebUSBDevice()
	deviceHandle, err := lowlevel.Open(device)
	if err != nil {
		return err
	}
	defer lowlevel.Close(deviceHandle)
	if _, err := lowlevel.Get_Descriptor(deviceHandle, lowlevel.DT_DEVICE, 0, []byte{}); err != nil {
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

// Reconnect tries to reconnect to find and reconnect to the WebUSB device
// of the wallet.
//
// If the wallet is not found it calls GetConfirm method to get a confirmation
// that it's required to try one more time.
func (trezor *trezorT) Reconnect() error {
	return routines.WebUSBReconnect(trezor)
}
