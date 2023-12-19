package drill

import (
	"encoding/json"
	"os/exec"
	"strings"
)


type PCIDevice struct {
	Slot        string `json:"slot"`
	Class       string `json:"class"`
	Vendor      string `json:"vendor"`
	Device      string `json:"device"`
	SVendor     string `json:"svendor"`
	SDevice     string `json:"sdevice"`
	Rev         string `json:"rev"`
	IOMMUGroup  string `json:"iommu_group"`
}


func ParsePCI(output string) ([]PCIDevice, error) {
	var devices []PCIDevice
	var device PCIDevice

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		keyValue := strings.SplitN(line, ":", 2)
		if len(keyValue) == 2 {
			key := strings.TrimSpace(keyValue[0])
			value := strings.TrimSpace(keyValue[1])

			switch key {
			case "Slot":
				if device.Slot != "" {
					devices = append(devices, device)
					device = PCIDevice{}
				}
				device.Slot = value
			case "Class":
				device.Class = value
			case "Vendor":
				device.Vendor = value
			case "Device":
				device.Device = value
			case "SVendor":
				device.SVendor = value
			case "SDevice":
				device.SDevice = value
			case "Rev":
				device.Rev = value
			case "IOMMUGroup":
				device.IOMMUGroup = value
			}
		}
	}
	if device.Slot != "" {
		devices = append(devices, device)
	}

	return devices, nil
}

func DrillPCI() (string, error) {
	out, err := exec.Command("lspci", "-vmm").Output()
	if err != nil {
		return "", err
	}

	devices, err := ParsePCI(string(out))
	if err != nil {
		return "", err
	}

	jsonData, err := json.MarshalIndent(devices, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
