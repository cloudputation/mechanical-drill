package transformer

import (
    "encoding/json"
    "fmt"
)

type Battery struct {
    State string `json:"state"`
    Percentage  string `json:"percentage"`
}

func GetBatteryDetails(jsonData string) {
    var batteries []Battery
    err := json.Unmarshal([]byte(jsonData), &batteries)
    if err != nil {
        fmt.Println("Error unmarshalling JSON:", err)
        return
    }

    var batteryDetails []interface{}
    for _, batteryDetail := range batteries {
        batteryDetails = append(batteryDetails, batteryDetail)
    }

    ExportDeviceDetails(batteryDetails, "battery")
}
