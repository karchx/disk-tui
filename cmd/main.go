package cmd

import (
	"bufio"
	"bytes"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

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
	dfPattern := regexp.MustCompile(`^\/dev\/sd[a-z]+\d*`)

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
			device := dfPattern.FindStringSubmatch(line)[0]

			//			rootPath := dfPattern.FindStringSubmatch(line)[2]
			log.Infof("Root: %s", device)
			//if ok := isMountDrive(device)
			//			isMountDrive(device)
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

// TODO: move pkg utils
var mu sync.Mutex

func concurrentRandom() int {
	mu.Lock()
	defer mu.Unlock()

	// Ensure that each goroutine gets a unique seed based on the current Unix timestamp.
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(101)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Concurrent Random Number:", concurrentRandom())
		}()
	}

	wg.Wait()
}
