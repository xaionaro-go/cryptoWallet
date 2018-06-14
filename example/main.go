package main

import (
	"github.com/xaionaro-go/cryptoWallet"
)

func main() {
	wallet := cryptoWallet.FindAny()
	if wallet == nil {
		panic("No wallets found")
	}

	err := wallet.Ping()
	if err != nil {
		panic(err)
	}
}
