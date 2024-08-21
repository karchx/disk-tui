package main

import (
	log "github.com/gothew/l-og"
	"github.com/karchx/disk-tui/cmd"
	"github.com/karchx/disk-tui/ui"
)

func main() {
	var cli cmd.Commands
	var tui ui.UI

	cli = cmd.NewCommand(cmd.Commands{
		Command: "sudo",
		Args:    []string{"fdisk", "-l"},
	})
	drives, err := cli.Drives()
	if err != nil {
		log.Error(err)
	}

	tui = ui.NewUI(ui.UI{
		Items: drives,
	})

	tui.Start()
	// cli = cmd.NewCommand(cmd.Commands{
	// 	Command: "sudo",
	// 	Args:    []string{"mount"},
	// 	Path:    "/mnt/setup-stiv",
	// })
	// message, err := cli.MountDisk(drives[0])
	// if err != nil {
	// 	log.Error(err)
	// }
	// fmt.Print(message)
}
