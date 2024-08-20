package main

import (
	"fmt"

	log "github.com/gothew/l-og"
	"github.com/karchx/disk-tui/cmd"
)

func main() {
	var cli cmd.Commands
	cli = cmd.NewCommand(cmd.Commands{
		Command: "sudo",
		Args:    []string{"fdisk", "-l"},
	})
	drives, err := cli.Drives()
	if err != nil {
		log.Error(err)
	}

	cli = cmd.NewCommand(cmd.Commands{
		Command: "sudo",
		Args:    []string{"mount"},
		Path:    "/mnt/setup-stiv",
	})
	message, err := cli.MountDisk(drives[0])
	if err != nil {
		log.Error(err)
	}
	fmt.Print(message)
}
