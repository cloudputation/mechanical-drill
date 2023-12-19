package drill

import (
	"os/exec"
)

func NetworkDrill() (string, error) {
	out, err := exec.Command("cat", "./tests/data/network.json").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
