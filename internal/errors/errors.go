package errors

import (
	"fmt"
)

var (
	// ErrNoWallet is returned when lost the connection with the wallet
	// device
	ErrNoWallet = fmt.Errorf("The wallet device is not found.")

	// ErrNotInitialized is returned when trying to do an action requires
	// an initialized device on a not initialized device
	ErrNotInitialized = fmt.Errorf("The wallet device is not initialized.")
)
