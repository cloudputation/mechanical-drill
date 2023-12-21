package engine

import (
    "github.com/cloudputation/mechanical-drill/packages/config"
)

type DrillScheduler struct {
    Network      int
    Storage      int
    CPU          int
}


func NewDrillScheduler() *DrillScheduler {
    return &DrillScheduler{
        Network: config.AppConfig.Drill.Network.Frequency,
        Storage: config.AppConfig.Drill.Storage.Frequency,
        CPU: config.AppConfig.Drill.CPU.Frequency,
    }
}
