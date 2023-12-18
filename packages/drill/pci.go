package drill

import (
  	"encoding/json"
  	"os/exec"
  	"strings"
)


type PCIDevice struct {
	Domain     string `json:"domain"`
	Bus        string `json:"bus"`
	Device     string `json:"device"`
	Function   string `json:"function"`
	Class      string `json:"class"`
	Vendor     string `json:"vendor"`
	DeviceName string `json:"device_name"`
	Revision   string `json:"revision"`
}


func ParsePCI(output string) ([]PCIDevice, error) {
	var devices []PCIDevice
	for _, line := range strings.Split(output, "\n") {
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			continue
		}

		// Extract the BDF (Bus/Device/Function)
		bdfParts := strings.Split(parts[0], ":")
		if len(bdfParts) < 2 {
			continue
		}
		busDevice := bdfParts[0]
		function := strings.TrimSuffix(bdfParts[1], ".")

		classVendorDevice := strings.Join(parts[1:], " ")
		splitIndex := strings.Index(classVendorDevice, ":")
		if splitIndex == -1 {
			continue
		}

		classInfo := classVendorDevice[:splitIndex]
		vendorDevice := classVendorDevice[splitIndex+2:]
		vendorDeviceParts := strings.SplitN(vendorDevice, " ", 2)
		if len(vendorDeviceParts) < 2 {
			continue
		}
		vendor := vendorDeviceParts[0]
		deviceName := vendorDeviceParts[1]

		devices = append(devices, PCIDevice{
			Domain:     "",
			Bus:        busDevice,
			Device:     strings.TrimSuffix(busDevice, "."),
			Function:   function,
			Class:      classInfo,
			Vendor:     vendor,
			DeviceName: deviceName,
			Revision:   "",
		})
	}
	return devices, nil
}

func DrillPCI() (string, error) {
    out, err := exec.Command("lspci").Output()
    if err != nil {
        return "", err
    }

    devices, err := ParsePCI(string(out))
    if err != nil {
        return "", err
    }

    deviceMap := make(map[string]PCIDevice)
    for _, device := range devices {
        deviceMap[device.DeviceName] = device
    }

    jsonData, err := json.MarshalIndent(deviceMap, "", "  ")
    if err != nil {
        return "", err
    }

    return string(jsonData), nil
}
