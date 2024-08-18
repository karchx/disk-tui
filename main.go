package main

import (
	"github.com/karchx/disk-tui/cmd"
)

func main() {
	cli := cmd.NewCommand()
	cli.Drives()
}
