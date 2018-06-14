package trezorBase

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/conejoninja/tesoro"
	"github.com/conejoninja/tesoro/pb/messages"
	"github.com/conejoninja/tesoro/transport"
	I "github.com/xaionaro-go/cryptoWallet/interfaces"
	routines "github.com/xaionaro-go/cryptoWallet/routines"
	base "github.com/xaionaro-go/cryptoWallet/wallets/satoshilabs"
	"github.com/zserge/hid"
)

type TrezorBase struct {
	base.SatoshilabsWalletBase
	Client tesoro.Client

	parent I.USBHIDWallet
}

func (trezor *TrezorBase) SetParent(parent I.USBHIDWallet) {
	if trezor.parent != nil {
		panic("trezor.parent is already defined")
	}
	trezor.parent = parent
}

func (trezor *TrezorBase) SetHIDDevice(device hid.Device) {
	var t transport.TransportHID
	t.SetDevice(device)
	trezor.Client.SetTransport(&t)
	trezor.USBHIDWalletBase.SetDevice(device)
}

func (trezor *TrezorBase) call(msg []byte) (string, uint16) {
	result, msgType := trezor.Client.Call(msg)

	switch messages.MessageType(msgType) {
	case messages.MessageType_MessageType_PinMatrixRequest:
		pin, err := trezor.GetPin("PIN", "", "Confirm", "Cancel")
		if err != nil {
			log.Print("Error", err)
		}
		result, msgType = trezor.call(trezor.Client.PinMatrixAck(string(pin)))

	case messages.MessageType_MessageType_ButtonRequest:

		result, msgType = trezor.call(trezor.Client.ButtonAck())

	case messages.MessageType_MessageType_PassphraseRequest:

		passphrase, err := trezor.GetPin("Passphrase", "", "Confirm", "Cancel")
		if err != nil {
			log.Print("Error", err)
		}
		result, msgType = trezor.call(trezor.Client.PassphraseAck(string(passphrase)))

	case messages.MessageType_MessageType_WordRequest:

		word, err := trezor.GetPin("Word", "", "OK", "Cancel")
		if err != nil {
			log.Print("Error", err)
		}
		result, msgType = trezor.call(trezor.Client.WordAck(string(word)))

	}

	return result, msgType
}

func (trezor *TrezorBase) ping(pingMsg string) (string, messages.MessageType) {
	pongMsg, msgType := trezor.Client.Call(trezor.Client.Ping(pingMsg, false, false, false))
	return pongMsg, messages.MessageType(msgType)
}

func (trezor *TrezorBase) Ping() error {
	if trezor.USBHIDWalletBase.GetDevice() == nil {
		return fmt.Errorf("trezor.USBHIDWalletBase.GetDevice() == nil")
	}
	if _, err := trezor.USBHIDWalletBase.GetDevice().HIDReport(); err != nil {
		return err
	}
	pongMsg, msgType := trezor.ping("ping")
	if pongMsg == "ping" {
		return nil
	}
	switch msgType {
	case messages.MessageType_MessageType_Success:
		return fmt.Errorf("The wallet device seems to be not initialized")
	}
	return fmt.Errorf("An unexpected behaviour of the wallet device: %v: %v", msgType, pongMsg)
}

func (trezor *TrezorBase) Reconnect() error {
	return routines.USBHIDReconnect(trezor.parent)
}

func (trezor *TrezorBase) CheckConnection() error {
	if trezor.Ping() == nil {
		return nil
	}

	return trezor.Reconnect()
}

// See https://github.com/satoshilabs/slips/blob/master/slip-0011.md
func (trezor *TrezorBase) CipherKeyValue(path string, isToEncrypt bool, keyName string, data, iv []byte, askOnEncode, askOnDecode bool) ([]byte, messages.MessageType) {
	result, msgType := trezor.call(trezor.Client.CipherKeyValue(isToEncrypt, keyName, data, tesoro.StringToBIP32Path(path), iv, askOnEncode, askOnDecode))
	return []byte(result), messages.MessageType(msgType)
}

func (trezor *TrezorBase) EncryptKey(path string, decryptedKey []byte, nonce []byte, trezorKeyname string) ([]byte, error) {
	// note: decryptedKey length should be aligned to 16 bytes

	trezor.CheckConnection()

	encryptedKey, msgTypeInt := trezor.CipherKeyValue(path, true, trezorKeyname, decryptedKey, nonce, false, true)

	msgType := messages.MessageType(msgTypeInt)
	switch msgType {
	case messages.MessageType_MessageType_Success, messages.MessageType_MessageType_CipheredKeyValue:
	case messages.MessageType_MessageType_Failure:
		return nil, fmt.Errorf(`Got an error from a trezor device: "%v" (the trezor device is busy?)`, string(encryptedKey))
	default:
		return nil, fmt.Errorf("Got an unexpected behaviour from a trezor device: %v: %v", msgType, string(encryptedKey))
	}

	return encryptedKey, nil
}

func (trezor *TrezorBase) DecryptKey(path string, encryptedKey []byte, nonce []byte, trezorKeyname string) ([]byte, error) {
	// note: encryptedKey length should be aligned to 16 bytes

	trezor.CheckConnection()

	// library "tesoro" requires hex-ed value for decryption
	encryptedKeyhexValue := hex.EncodeToString(encryptedKey)
	if len(encryptedKeyhexValue)%2 != 0 {
		log.Panic(`len(hexValue) is odd`)
	}
	for len(encryptedKeyhexValue)%32 != 0 {
		encryptedKeyhexValue += "00"
	}

	decryptedKey, msgType := trezor.CipherKeyValue(path, false, trezorKeyname, []byte(encryptedKeyhexValue), nonce, false, true)

	switch msgType {
	case messages.MessageType_MessageType_Success, messages.MessageType_MessageType_CipheredKeyValue:
	case messages.MessageType_MessageType_Failure:
		return nil, fmt.Errorf(`Got an error from a trezor device: %v (the trezor device is busy?)`, string(decryptedKey)) // if an error occurs then the error description is returned into "decryptedKey" as a string
	default:
		return nil, fmt.Errorf("Got an unexpected behaviour from a trezor device: %v: %v", msgType, string(encryptedKey))
	}

	return decryptedKey, nil
}
