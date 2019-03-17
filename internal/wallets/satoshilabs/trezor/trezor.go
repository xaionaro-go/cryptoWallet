package trezorBase

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"github.com/conejoninja/tesoro"
	"github.com/conejoninja/tesoro/pb/messages"
	I "github.com/xaionaro-go/cryptoWallet/interfaces"
	"github.com/xaionaro-go/cryptoWallet/internal/errors"
	"github.com/xaionaro-go/cryptoWallet/internal/wallets/satoshilabs"
)

// TrezorBase is an implementation of common properties and methods of
// all trezor devices
type TrezorBase struct {
	wallet I.Wallet // The parent
	satoshilabsWallet.Base
	Client tesoro.Client

	defaultAskOnEncode bool
}

// SetName sets a new name of the device
func (trezor *TrezorBase) SetWallet(wallet I.Wallet) {
	if trezor.wallet != nil {
		panic("trezor.wallet is already not nil")
	}
	trezor.wallet = wallet
}

// SetDefaultAskOnEncode is a method to set the value which will be passed
// as argument "askOnEncode" to "CipherKeyValue" from method "EncryptKey"
func (trezor *TrezorBase) SetDefaultAskOnEncode(newDefaultAskOnEncode bool) {
	trezor.defaultAskOnEncode = newDefaultAskOnEncode
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

func (trezor *TrezorBase) Ping(pingMsg string) (string, messages.MessageType) {
	pongMsg, msgType := trezor.Client.Call(trezor.Client.Ping(pingMsg, false, false, false))
	return pongMsg, messages.MessageType(msgType)
}

type initializeResponse struct {
	Vendor          string `json:"vendor"`
	MajorVersion    int    `json:"major_version"`
	MinorVersion    int    `json:"minor_version"`
	PatchVersion    int    `json:"patch_version"`
	BootloaderMode  bool   `json:"bootloader_mode"`
	FirmwarePresent bool   `json:"firmware_present"`
}

// Reset sends an empty initialize package and checkes if the response
// is correct. This function resets the state of the device and checks
// if it is initialized. If the device is not initialized then
// ErrNotInitialized will be returned.
//
// See also: https://doc.satoshilabs.com/trezor-tech/api-workflows.html#initialize-features
func (trezor *TrezorBase) Reset() error {
	str, msgTypeRaw := trezor.call(trezor.Client.Initialize())
	msgType := messages.MessageType(msgTypeRaw)
	if msgType != messages.MessageType_MessageType_Features {
		return fmt.Errorf("Got an unexpected behaviour from a trezor device: %v: %v", msgType, str)
	}

	// An example of not the answer of a not initialized device:
	//
	// {"vendor":"bitcointrezor.com","major_version":1,"minor_version":4,"patch_version":0,"bootloader_mode":true,"firmware_present":false}
	var response initializeResponse
	err := json.Unmarshal([]byte(str), &response)
	if err != nil {
		return err
	}

	if response.BootloaderMode {
		return errors.ErrNotInitialized
	}

	return nil
}

// CheckConnection checks the connection to the device and reconnects if required
func (trezor *TrezorBase) CheckConnection() error {
	if trezor.wallet.Ping() == nil {
		return nil
	}

	return trezor.wallet.Reconnect()
}

// CipherKeyValue is a function of symmetric encryption of key-value paris
// using deterministic hierarchy
//
// See https://github.com/satoshilabs/slips/blob/master/slip-0011.md
func (trezor *TrezorBase) CipherKeyValue(path string, isToEncrypt bool, keyName string, data, iv []byte, askOnEncode, askOnDecode bool) ([]byte, messages.MessageType) {
	result, msgType := trezor.call(trezor.Client.CipherKeyValue(isToEncrypt, keyName, data, tesoro.StringToBIP32Path(path), iv, askOnEncode, askOnDecode))
	return []byte(result), messages.MessageType(msgType)
}

// EncryptKey encrypts a key using a symmetric algorithm. The key length
// should be a multiple of 16 bytes.
//
// - `path` is a BIP32 path;
//
// - `decryptedKey` is a key to be encrypted;
//
// - `nonce` is optional "number that can only be used once",
//    see https://en.wikipedia.org/wiki/Cryptographic_nonce;
//
// - `trezorKeyname` is a key name that affects on encrypts and displays
//    on the screen of a trezor device.
func (trezor *TrezorBase) EncryptKey(path string, decryptedKey []byte, nonce []byte, trezorKeyname string) ([]byte, error) {
	// note: decryptedKey length should be aligned to 16 bytes

	trezor.CheckConnection()

	encryptedKey, msgTypeInt := trezor.CipherKeyValue(path, true, trezorKeyname, decryptedKey, nonce, trezor.defaultAskOnEncode, true)

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

// DecryptKey decrypts a key using a symmetric algorithm. The key length
// should be a multiple of 16 bytes.
//
// - `path` is a BIP32 path;
//
// - `encryptedKey` is a key to be decrypted;
//
// - `nonce` is "number that can only be used once",
//    see https://en.wikipedia.org/wiki/Cryptographic_nonce;
//
// - `trezorKeyname` is a key name that affects on encrypts and displays
//    on the screen of a trezor device.
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

	decryptedKey, msgType := trezor.CipherKeyValue(path, false, trezorKeyname, []byte(encryptedKeyhexValue), nonce, trezor.defaultAskOnEncode, true)

	switch msgType {
	case messages.MessageType_MessageType_Success, messages.MessageType_MessageType_CipheredKeyValue:
	case messages.MessageType_MessageType_Failure:
		return nil, fmt.Errorf(`Got an error from a trezor device: %v (the trezor device is busy?)`, string(decryptedKey)) // if an error occurs then the error description is returned into "decryptedKey" as a string
	default:
		return nil, fmt.Errorf("Got an unexpected behaviour from a trezor device: %v: %v", msgType, string(encryptedKey))
	}

	return decryptedKey, nil
}
