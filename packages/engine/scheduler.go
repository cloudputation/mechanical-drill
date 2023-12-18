package engine

import (
    "github.com/cloudputation/mechanical-drill/packages/config"
)


type DrillScheduler struct {
  PCI int
}


func NewDrillScheduler() *DrillScheduler {
    return &DrillScheduler{
        PCI: config.AppConfig.Drill.PCI.Frequency,
    }
}
