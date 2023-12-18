package engine

import (
	"encoding/json"
	"fmt"
	"os/exec"
  "reflect"
  "regexp"
  "strconv"

  "github.com/cloudputation/mechanical-drill/packages/drill"
)


func PrintDeviceDetails(jsonData string) {
    var devicesMap map[string]drill.PCIDevice
    err := json.Unmarshal([]byte(jsonData), &devicesMap)
    if err != nil {
        fmt.Println("Error parsing JSON:", err)
        return
    }

    var i int = 1
    for deviceName, device := range devicesMap {
        unwantedChars := regexp.MustCompile(`[ \-.,#/[\]()]+`)
        cleanedDeviceName := unwantedChars.ReplaceAllString(deviceName, "_")

        fields := []string{"Domain", "Bus", "Device", "Function", "Class", "Vendor", "DeviceName", "Revision"}
        fmt.Printf("\n\n")
        fmt.Printf("%s\n", cleanedDeviceName)
        for _, field := range fields {
            fieldValue := GetFieldValue(device, field)
            fmt.Printf("\t%s: %s\n", field, fieldValue)

            index := strconv.Itoa(i)
            kv := fmt.Sprintf("md.pci.device%s.%s=%s", index, field, fieldValue)
            cmd := exec.Command("nomad", "node", "meta", "apply", kv)

            if err := cmd.Run(); err != nil {
                fmt.Println("Error executing command:", err)
            }
        }
        i++
    }
}

func GetFieldValue(device drill.PCIDevice, field string) string {
	val := reflect.ValueOf(device)
	fieldVal := val.FieldByName(field)
	if !fieldVal.IsValid() {
	   return ""
	}
	return fieldVal.String()
}
