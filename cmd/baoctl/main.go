package main

import (
	_ "baoctl/bootstrap"
	"baoctl/pkg/util"
	"os"

	"baoctl/pkg/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		util.PrintError(err)
		os.Exit(1)
	}
}
