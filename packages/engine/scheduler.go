package engine

import (
    "github.com/cloudputation/mechanical-drill/packages/config"
)

type DrillScheduler struct {
    System        int
    Battery       int
    Bridge        int
    Memory        int
    Processor     int
    Address       int
    Storage       int
    Disk          int
    Tape          int
    Bus           int
    Network       int
    Display       int
    Input         int
    Printer       int
    Multimedia    int
    Communication int
    Power         int
    Volume        int
    Generic       int
}

func NewDrillScheduler() *DrillScheduler {
    ds := &DrillScheduler{}

    if config.AppConfig.Drill.System != nil {
        ds.System = config.AppConfig.Drill.System.Frequency
    }
    if config.AppConfig.Drill.Battery != nil {
        ds.Battery = config.AppConfig.Drill.Battery.Frequency
    }
    if config.AppConfig.Drill.Bridge != nil {
        ds.Bridge = config.AppConfig.Drill.Bridge.Frequency
    }
    if config.AppConfig.Drill.Memory != nil {
        ds.Memory = config.AppConfig.Drill.Memory.Frequency
    }
    if config.AppConfig.Drill.Processor != nil {
        ds.Processor = config.AppConfig.Drill.Processor.Frequency
    }
    if config.AppConfig.Drill.Address != nil {
        ds.Address = config.AppConfig.Drill.Address.Frequency
    }
    if config.AppConfig.Drill.Storage != nil {
        ds.Storage = config.AppConfig.Drill.Storage.Frequency
    }
    if config.AppConfig.Drill.Disk != nil {
        ds.Disk = config.AppConfig.Drill.Disk.Frequency
    }
    if config.AppConfig.Drill.Tape != nil {
        ds.Tape = config.AppConfig.Drill.Tape.Frequency
    }
    if config.AppConfig.Drill.Bus != nil {
        ds.Bus = config.AppConfig.Drill.Bus.Frequency
    }
    if config.AppConfig.Drill.Network != nil {
        ds.Network = config.AppConfig.Drill.Network.Frequency
    }
    if config.AppConfig.Drill.Display != nil {
        ds.Display = config.AppConfig.Drill.Display.Frequency
    }
    if config.AppConfig.Drill.Input != nil {
        ds.Input = config.AppConfig.Drill.Input.Frequency
    }
    if config.AppConfig.Drill.Printer != nil {
        ds.Printer = config.AppConfig.Drill.Printer.Frequency
    }
    if config.AppConfig.Drill.Multimedia != nil {
        ds.Multimedia = config.AppConfig.Drill.Multimedia.Frequency
    }
    if config.AppConfig.Drill.Communication != nil {
        ds.Communication = config.AppConfig.Drill.Communication.Frequency
    }
    if config.AppConfig.Drill.Power != nil {
        ds.Power = config.AppConfig.Drill.Power.Frequency
    }
    if config.AppConfig.Drill.Volume != nil {
        ds.Volume = config.AppConfig.Drill.Volume.Frequency
    }
    if config.AppConfig.Drill.Generic != nil {
        ds.Generic = config.AppConfig.Drill.Generic.Frequency
    }


    return ds
}
