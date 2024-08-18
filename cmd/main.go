package cmd

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"regexp"
	"strings"

	log "github.com/gothew/l-og"
)

type Commands struct {
	command string
	args    []string
}

func NewCommand() Commands {
	return Commands{
		command: "fdisk",
		args:    []string{"-l"},
	}
}

func (c *Commands) Drives() string {
	//var drives []string
	//	driveMap := make(map[string]bool)
	dfPattern := regexp.MustCompile("^(\\/[^ ]+)[^%]+%[ ]+(.+)$")

	cmd := c.command
	//	args := c.args
	cm := exec.Command("sudo", "fdisk", "-l")
	cm.Stderr = os.Stderr
	cm.Stdin = os.Stdin

	out, err := cm.Output()

	if err != nil {
		log.Errorf("Error calling %s: %s", cmd, err)
	}

	s := bufio.NewScanner(bytes.NewReader(out))

	for s.Scan() {
		line := s.Text()
		if dfPattern.MatchString(line) {
			device := dfPattern.FindStringSubmatch(line)[1]
			log.Infof("Device: %s", device)
			//rootPath := dfPattern.FindStringSubmatch(line)[2]
			//if ok := isMountDrive(device)
			//isMountDrive(device)
			//log.Infof("Device: %s is: %s", device, isMountDrive(device))
		}
	}

	output := string(out[:])
	return output
}

func isMountDrive(device string) bool {
	deviceVerifier := "ID_USB_DRIVER=usb-storage"
	cmd := "udevadm" // Command default, no changes
	args := []string{"info", "-q", "property", "-n", device}
	out, err := exec.Command(cmd, args...).Output()

	if err != nil {
		log.Errorf("Error checking device %s %s", device, err)
	}

	//	log.Infof("Device: %s return: %s", device, string(out))

	if strings.Contains(string(out), deviceVerifier) {
		return true
	}

	return false
}
