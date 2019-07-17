package main

import (
	"os"
	"github.com/sunnyRK/Gochain/cli"
	// "github.com/sunnyRK/Gochain/wallet"
)


func main() {
	defer os.Exit(0)
	cli := cli.CommandLine{}
	cli.Run()

	// w := wallet.MakeWallet()
	// w.Address()
}