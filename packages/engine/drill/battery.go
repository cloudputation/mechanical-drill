package drill

import (
	"encoding/json"
	"os/exec"
	"regexp"
	"strings"
)

func GetBatteryStatus() (string, error) {
	cmd := "upower -i $(upower -e | grep 'BAT') | grep -E 'state|percentage'"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", err
	}

	batteryStatus, err := parseBatteryStatus(string(out))
	if err != nil {
		return "", err
	}

	jsonData, err := json.Marshal(batteryStatus)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func parseBatteryStatus(output string) ([]map[string]string, error) {
	lines := strings.Split(output, "\n")
	batteryStatus := make([]map[string]string, 0)

	statusMap := make(map[string]string)
	for _, line := range lines {
		re := regexp.MustCompile(`^\s*(state|percentage):\s*(.*)$`)
		matches := re.FindStringSubmatch(line)
		if len(matches) == 3 {
			statusMap[matches[1]] = strings.TrimSpace(matches[2])
		}
	}

	if len(statusMap) > 0 {
		batteryStatus = append(batteryStatus, statusMap)
	}

	return batteryStatus, nil
}
