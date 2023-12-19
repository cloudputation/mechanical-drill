package engine

import (
    "fmt"
    "time"

    "github.com/cloudputation/mechanical-drill/packages/config"
    "github.com/cloudputation/mechanical-drill/packages/engine/drill"
    "github.com/cloudputation/mechanical-drill/packages/transformer"
)

func (ds *DrillScheduler) RunController() {
    networkTicker := time.NewTicker(time.Duration(ds.Network) * time.Second)
    storageTicker := time.NewTicker(time.Duration(ds.Storage) * time.Second)

    if config.AppConfig.Drill.Network.Enabled {
        ds.runNetworkDrill()
    }
    if config.AppConfig.Drill.Storage.Enabled {
        ds.runStorageDrill()
    }

    go func() {
        for {
            select {
            case <-networkTicker.C:
                if config.AppConfig.Drill.Network.Enabled {
                    ds.runNetworkDrill()
                }
            case <-storageTicker.C:
                if config.AppConfig.Drill.Storage.Enabled {
                    ds.runStorageDrill()
                }
            }
        }
    }()
    select {}
}

func (ds *DrillScheduler) runNetworkDrill() {
    jsonData, err := drill.NetworkDrill()
    if err != nil {
        fmt.Println("Network Drill Error:", err)
        return
    }
    transformer.ExportNetworkDetails(jsonData)
}

func (ds *DrillScheduler) runStorageDrill() {
    jsonData, err := drill.StorageDrill()
    if err != nil {
        fmt.Println("Storage Drill Error:", err)
        return
    }
    transformer.ExportStorageDetails(jsonData)
}
