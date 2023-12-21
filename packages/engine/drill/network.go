package drill

import (
	"os/exec"
)

func NetworkDrill() (string, error) {
	out, err := exec.Command("lshw", "-json", "-class", "network").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
