package transformer

import (
    "encoding/json"
    "fmt"
)

type StorageDevice struct {
    Id             string                 `json:"id"`
    Class          string                 `json:"class"`
    Description    string                 `json:"description"`
    Product        string                 `json:"product"`
    Vendor         string                 `json:"vendor"`
    PhysId         string                 `json:"physid"`
    BusInfo        string                 `json:"businfo"`
    LogicalName    string                 `json:"logicalname"`
    Version        string                 `json:"version"`
    Serial         string                 `json:"serial"`
    Width          int                    `json:"width"`
    Clock          int                    `json:"clock"`
    Configuration  map[string]interface{} `json:"configuration"`
    Capabilities   map[string]interface{} `json:"capabilities"`
}

func ExportStorageDetails(jsonData string) {
    var storageDevices []StorageDevice
    err := json.Unmarshal([]byte(jsonData), &storageDevices)
    if err != nil {
        fmt.Println("Error unmarshalling JSON:", err)
        return
    }

    var devices []interface{}
    for _, device := range storageDevices {
        devices = append(devices, device)
    }

    ExportDeviceDetails(devices)
}
