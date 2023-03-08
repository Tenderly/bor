package main

import (
	"os"

	"github.com/tenderly/bor/go-ethereum/internal/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
