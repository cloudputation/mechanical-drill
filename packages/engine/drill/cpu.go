package drill

import (
	"os/exec"
)

func CPUDrill() (string, error) {
	out, err := exec.Command("lshw", "-json", "-class", "processor").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
