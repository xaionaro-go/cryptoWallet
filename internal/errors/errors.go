package errors

import (
	"fmt"
)

var (
	// ErrNoWallet is returned when lost the connection with the wallet device
	ErrNoWallet = fmt.Errorf("The wallet device is not found.")
)
