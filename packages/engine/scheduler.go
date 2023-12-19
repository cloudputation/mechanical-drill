package engine

import (
    "sync"

    "github.com/cloudputation/mechanical-drill/packages/config"
)

type DrillScheduler struct {
    Network      int
    Storage      int
    mutex        sync.Mutex
    networkBusy  bool
    storageBusy  bool
}


func NewDrillScheduler() *DrillScheduler {
    return &DrillScheduler{
        Network: config.AppConfig.Drill.Network.Frequency,
        Storage: config.AppConfig.Drill.Storage.Frequency,
    }
}
