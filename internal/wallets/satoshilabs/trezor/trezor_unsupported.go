// +build !cgo iso !linux,!darwin,!windows

package trezorBase

import (
	"log"

	"github.com/conejoninja/tesoro/pb/messages"
	"github.com/xaionaro-go/cryptoWallet/internal/errors"
	"github.com/xaionaro-go/cryptoWallet/internal/wallets/satoshilabs"
	"github.com/zserge/hid"
)

// TrezorBase -- see trezor.go
type TrezorBase struct {
	satoshilabsWallet.Base
}

// SetHIDDevice -- see trezor.go
func (trezor *TrezorBase) SetHIDDevice(device hid.Device) {
	log.Panic(errors.ErrNotSupportedPlatform)
}
// Reset -- see trezor.go
func (trezor *TrezorBase) Reset() error {
	return errors.ErrNotSupportedPlatform
}
// Ping -- see trezor.go
func (trezor *TrezorBase) Ping() error {
	return errors.ErrNotSupportedPlatform
}
// Reconnect -- see trezor.go
func (trezor *TrezorBase) Reconnect() error {
	return errors.ErrNotSupportedPlatform
}
// CheckConnection -- see trezor.go
func (trezor *TrezorBase) CheckConnection() error {
	return errors.ErrNotSupportedPlatform
}
// CipherKeyValue -- see trezor.go
func (trezor *TrezorBase) CipherKeyValue(path string, isToEncrypt bool, keyName string, data, iv []byte, askOnEncode, askOnDecode bool) ([]byte, messages.MessageType) {
	log.Panic(errors.ErrNotSupportedPlatform)
	return nil, messages.MessageType_MessageType_Failure
}
// EncryptKey -- see trezor.go
func (trezor *TrezorBase) EncryptKey(path string, decryptedKey []byte, nonce []byte, trezorKeyname string) ([]byte, error) {
	return nil, errors.ErrNotSupportedPlatform
}
// DecryptKey -- see trezor.go
func (trezor *TrezorBase) DecryptKey(path string, encryptedKey []byte, nonce []byte, trezorKeyname string) ([]byte, error) {
	return nil, errors.ErrNotSupportedPlatform
}
