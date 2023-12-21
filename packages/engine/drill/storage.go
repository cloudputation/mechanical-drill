package drill

import (
	"os/exec"
)

func StorageDrill() (string, error) {
	out, err := exec.Command("lshw", "-json", "-class", "storage").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
