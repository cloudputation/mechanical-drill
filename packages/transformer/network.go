package transformer

import (
    "encoding/json"
    "fmt"
)

type NetworkDevice struct {
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
  	Units          string                 `json:"units"`
  	Size           int64                  `json:"size"`
  	Capacity       int64                  `json:"capacity"`
  	Width          int                    `json:"width"`
  	Clock          int                    `json:"clock"`
  	Configuration  map[string]interface{} `json:"configuration"`
  	Capabilities   map[string]interface{} `json:"capabilities"`
}

func ExportNetworkDetails(jsonData string) {
    var storageDevices []NetworkDevice
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
