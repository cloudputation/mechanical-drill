package transformer

import (
    "encoding/json"
    "fmt"
)

type ProcessorDevice struct {
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

func ExportCPUDetails(jsonData string) {
    var processorDevices []ProcessorDevice
    err := json.Unmarshal([]byte(jsonData), &processorDevices)
    if err != nil {
        fmt.Println("Error unmarshalling JSON:", err)
        return
    }

    var devices []interface{}
    for _, device := range processorDevices {
        devices = append(devices, device)
    }

    ExportDeviceDetails(devices)
}
