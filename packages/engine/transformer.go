package engine

import (
    "encoding/json"
    "fmt"
    "os/exec"
    "reflect"
    "strconv"

    "github.com/cloudputation/mechanical-drill/packages/engine/drill"
)

func PrintDeviceDetails(jsonData string) {
    var devices []drill.PCIDevice
    err := json.Unmarshal([]byte(jsonData), &devices)
    if err != nil {
        fmt.Println("Error parsing JSON:", err)
        return
    }

    var i int = 1
    for _, device := range devices {
        fields := []string{"Slot", "Class", "Vendor", "Device", "SVendor", "SDevice", "Rev", "IOMMUGroup"}
        fmt.Printf("\n\n")
        fmt.Printf("%s\n", device)
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
    if !fieldVal.IsValid() || fieldVal.String() == "" {
        return "N/A"
    }
    return fieldVal.String()
}
