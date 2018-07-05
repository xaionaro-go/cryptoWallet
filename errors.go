package cryptoWallet

import (
	"github.com/xaionaro-go/cryptoWallet/internal/errors"
)

var (
	// ErrNoWallet is returned when lost the connection with the wallet
	// device
	ErrNoWallet = errors.ErrNoWallet

	// ErrNotInitialized is returned when trying to do an action requires
	// an initialized device on a not initialized device
	ErrNotInitialized = errors.ErrNotInitialized

	// ErrNotSupportedPlatform is returned/paniced when trying use this library
	// on an unsupported platform
	ErrNotSupportedPlatform = errors.ErrNotSupportedPlatform
)
