package cryptoWalletInterfaces

import (
	"github.com/conejoninja/tesoro/pb/messages"
)

// Trezor is an interface that publishes methods of devices "Trezor One" and "Trezor T"
type Trezor interface {
	USBWallet

	// CipherKeyValue is a function of symmetric encryption of key-value paris
	// using deterministic hierarchy
	//
	// See https://github.com/satoshilabs/slips/blob/master/slip-0011.md
	CipherKeyValue(path string, isToEncrypt bool, keyName string, data, iv []byte, askOnEncode, askOnDecode bool) ([]byte, messages.MessageType)

	// SetDefaultAskOnEncode is a method to set the value which will be passed
	// as argument "askOnEncode" to "CipherKeyValue" from method "EncryptKey"
	SetDefaultAskOnEncode(newDefaultAskOnEncode bool)
}
