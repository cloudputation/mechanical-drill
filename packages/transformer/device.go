package transformer

import (
    "encoding/json"
    "fmt"
)

type Device struct {
    Id            string                 `json:"id"`
    Class         string                 `json:"class"`
    Description   string                 `json:"description"`
    Product       string                 `json:"product"`
    Vendor        string                 `json:"vendor"`
    PhysId        string                 `json:"physid"`
    BusInfo       string                 `json:"businfo"`
    Version       string                 `json:"version"`
    Slot          string                 `json:"slot"`
    Units         string                 `json:"units"`
    Size          int64                  `json:"size"`
    Capacity      int64                  `json:"capacity"`
    Width         int                    `json:"width"`
    Clock         int                    `json:"clock"`
    Configuration map[string]interface{} `json:"configuration"`
    Capabilities  map[string]interface{} `json:"capabilities"`
}

func GetDeviceDetails(jsonData string) {
    var devices []Device
    err := json.Unmarshal([]byte(jsonData), &devices)
    if err != nil {
        fmt.Println("Error unmarshalling JSON:", err)
        return
    }

    var devicesDetails []interface{}
    for _, deviceDetail := range devices {
        devicesDetails = append(devicesDetails, deviceDetail)
    }

    ExportDeviceDetails(devicesDetails)
}
