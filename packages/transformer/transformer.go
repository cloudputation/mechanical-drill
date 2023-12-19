package transformer

import (
    "fmt"
    "os/exec"
    "reflect"
    "strconv"
    "strings"
)

func ExportDeviceDetails(devices []interface{}) {
    for i, device := range devices {
        processDeviceFields(device, strconv.Itoa(i))
    }
}

func processDeviceFields(device interface{}, deviceIndex string) {
    val := reflect.ValueOf(device)
    deviceType := val.Type()

    classField := val.FieldByName("Class")
    if !classField.IsValid() {
        fmt.Println("Class field not found in the struct")
        return
    }
    classValue := strings.ToLower(classField.String())

    for i := 0; i < val.NumField(); i++ {
        field := deviceType.Field(i)
        fieldValue := val.Field(i).Interface()

        if field.Type.Kind() == reflect.Map {
            processMapField(field.Name, fieldValue, deviceIndex, classValue)
            continue
        }

        kv := fmt.Sprintf("md.%s.device%s.%s=%v", classValue, deviceIndex, field.Name, fieldValue)
        fmt.Printf("[INFO] created hardware entry: %s\n", kv)
        applyNomadMetadata(kv)
    }
}

func processMapField(fieldName string, value interface{}, deviceIndex string, deviceClass string) {
    m, ok := value.(map[string]interface{})
    if !ok {
        fmt.Printf("Error processing map field: %s\n", fieldName)
        return
    }

    for k, v := range m {
        k = strings.ReplaceAll(k, ".", "")

        if len(k) > 0 && k[0] >= '0' && k[0] <= '9' {
            k = "d" + k
        }

        kv := fmt.Sprintf("md.%s.device%s.%s.%s=%v", deviceClass, deviceIndex, strings.ToLower(fieldName), k, v)
        fmt.Printf("[INFO] created hardware entry: %s\n", kv)
        applyNomadMetadata(kv)
    }
}



func applyNomadMetadata(kv string) {
    cmd := exec.Command("nomad", "node", "meta", "apply", kv)
    if err := cmd.Run(); err != nil {
        fmt.Println("Error executing command:", err)
    }
}
