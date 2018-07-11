// +build linux

package cryptoWalletRoutines

import (
	"fmt"
	"log"

	"github.com/conejoninja/hid"
	I "github.com/xaionaro-go/cryptoWallet/interfaces"
	"github.com/xaionaro-go/cryptoWallet/internal/errors"
)

// USBHIDReconnect tries to reconnect to find and reconnect to the
// USB HID device of the wallet `parent`.
//
// If the wallet is not found it calls GetConfirm method of the `parent` to
// get a confirmation that it's required to try one more time.
func USBHIDReconnect(parent I.USBHIDWallet) error {
	success := false
	for !success {
		hid.UsbWalk(func(device hid.Device) {
			info := device.Info()
			if info.Vendor == parent.GetVendorID() && info.Product == parent.GetProductID() && info.Interface == parent.GetInterfaceID() {
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
				return errors.ErrNoWallet
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
