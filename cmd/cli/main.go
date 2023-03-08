package main

import (
	"os"

	"github.com/tenderly/bor/internal/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
