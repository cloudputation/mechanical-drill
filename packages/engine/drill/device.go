package drill

import (
	"os/exec"
)

func Drill(deviceClass string) (string, error) {
	out, err := exec.Command("lshw", "-json", "-class", deviceClass).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
