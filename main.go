package main

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	log "github.com/gothew/l-og"
	"github.com/karchx/disk-tui/cmd"
	"github.com/karchx/disk-tui/ui"
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

	if _, err := tea.NewProgram(ui.NewModel(drives), tea.WithAltScreen()).Run(); err != nil {
		log.Errorf("Error running program: %s \n", err)
		os.Exit(1)
	}

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
