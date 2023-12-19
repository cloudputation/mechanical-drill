package drill

import (
	"os/exec"
)

func StorageDrill() (string, error) {
	out, err := exec.Command("cat", "./tests/data/storage.json").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
