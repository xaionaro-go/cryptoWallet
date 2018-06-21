// +build linux

package cryptoWalletRoutines

import (
	"fmt"
	"log"

	I "github.com/xaionaro-go/cryptoWallet/internal/interfaces"
	"github.com/zserge/hid"
)

var (
	ErrNoWallet = fmt.Errorf("The wallet device is not found.")
)

func USBHIDReconnect(parent I.USBHIDWallet) error {
	success := false
	for !success {
		hid.UsbWalk(func(device hid.Device) {
			info := device.Info()
			if info.Vendor == parent.GetVendorId() && info.Product == parent.GetProductId() && info.Interface == parent.GetInterfaceId() {
				parent.SetHIDDevice(device)
				success = true
				return
			}
		})
		if !success {
			log.Printf("No %v devices found.", parent.Name())
			shouldContinue, err := parent.GetConfirm(fmt.Sprintf("The %v device is not found", parent.Name()), "Please check the connection to the device", "Retry", "Cancel")
			if err != nil {
				return err
			}
			if !shouldContinue {
				return ErrNoWallet
			}
		} else {
			err := parent.Ping()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
