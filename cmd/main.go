package cmd

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	log "github.com/gothew/l-og"
)

type Commands struct {
	Command string
	Args    []string
	Path    string
}

func NewCommand(command Commands) Commands {
	return command
}

func (c *Commands) Drives() ([]string, error) {
	var drives []string

	driveMap := make(map[string]bool)

	// Regex for filter /dev/sdx
	dfPattern := regexp.MustCompile(`^\/dev\/sd[a-z]+\d*`)

	cmd := c.Command
	args := c.Args
	cm := exec.Command(cmd, args...)
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
			device := dfPattern.FindStringSubmatch(line)[0]
			for key, ok := range isMountDrive(device) {
				if ok {
					keyDevice := device + " - " + key // Name device + type device
					driveMap[keyDevice] = true
				}
			}
		}
	}

	for k := range driveMap {
		drives = append(drives, k)
	}

	return drives, nil
}

func (c *Commands) MountDisk(drive string) (string, error) {
	if c.Path == "" {
		return "", errors.New("Missing path for mount drive")
	}

	cmd := c.Command
	args := c.Args
	args = append(args, drive)
	args = append(args, c.Path)
	cm := exec.Command(cmd, args...)

	_, err := cm.Output()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s is mount", c.Path), nil
}

// TODO: move to utils pkg
func isMountDrive(device string) map[string]bool {
	deviceMap := make(map[string]bool)
	validDevice := "ID_USB_DRIVER=uas|ID_USB_DRIVER=usb-storage"
	deviceVerifier := strings.Split(validDevice, "|")
	cmd := "udevadm" // Command default, no changes
	args := []string{"info", "-q", "property", "-n", device}
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		log.Errorf("Error checking device %s %s", device, err)
	}

	if strings.Contains(string(out), deviceVerifier[0]) || strings.Contains(string(out), deviceVerifier[1]) {
		infoDriver := getInfoDriveLine(string(out), "ID_USB_DRIVER", "=")
		deviceMap[infoDriver] = true
		return deviceMap
	}
	deviceMap[""] = false
	return deviceMap
}

func getInfoDriveLine(info string, keyFilter string, separator string) string {
	var findLine string
	scanner := bufio.NewScanner(strings.NewReader(info))

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyFilter) {
			findLine = line
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Errorf("Error leyendo el texto: %s\n", err)
	}
	splitFind := strings.Split(findLine, separator)

	// get last value and return after
	findLine = splitFind[len(splitFind)-1]

	// uas is 'hard drive'
	if findLine == "uas" {
		findLine = "hard-drive"
	} else {
		findLine = findLine
	}

	return findLine
}
