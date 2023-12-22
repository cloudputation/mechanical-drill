package engine

import (
    "fmt"
    "reflect"
    "strings"
    "time"

    "github.com/cloudputation/mechanical-drill/packages/config"
    "github.com/cloudputation/mechanical-drill/packages/engine/drill"
    "github.com/cloudputation/mechanical-drill/packages/transformer"
)

var drillHeads = []string{
    "system",
    "bridge",
    "battery",
    "memory",
    "processor",
    "address",
    "storage",
    "disk",
    "tape",
    "bus",
    "network",
    "display",
    "input",
    "printer",
    "multimedia",
    "communication",
    "power",
    "volume",
    "generic",
}


func (ds *DrillScheduler) RunController() {
    tickers := make(map[string]*time.Ticker)

    for _, drillHead := range drillHeads {
        drillConfigVal := reflect.ValueOf(&config.AppConfig.Drill).Elem()
        drillConfigField := drillConfigVal.FieldByName(strings.Title(drillHead))
        if drillConfigField.IsValid() && !drillConfigField.IsNil() {
            frequency := int(drillConfigField.Elem().FieldByName("Frequency").Int())
            tickers[drillHead] = time.NewTicker(time.Duration(frequency) * time.Second)
        }
    }

    for _, drillHead := range drillHeads {
        if drillConfig := getDrillConfig(&config.AppConfig.Drill, drillHead); drillConfig {
          go func(dt string) {
              ds.runDrill(dt)
          }(drillHead)
        }
    }


     for {
         select {
         case <-safeReceive(tickers["system"]):
             if config.AppConfig.Drill.System != nil && config.AppConfig.Drill.System.Enabled {
                 go func() { ds.runDrill("system") }()
             }
         case <-safeReceive(tickers["battery"]):
             if config.AppConfig.Drill.Battery != nil && config.AppConfig.Drill.Battery.Enabled {
                 go func() { ds.runDrill("battery") }()
             }
         case <-safeReceive(tickers["bridge"]):
             if config.AppConfig.Drill.Bridge != nil && config.AppConfig.Drill.Bridge.Enabled {
                 go func() { ds.runDrill("bridge") }()
             }
         case <-safeReceive(tickers["memory"]):
             if config.AppConfig.Drill.Memory != nil && config.AppConfig.Drill.Memory.Enabled {
                 go func() { ds.runDrill("memory") }()
             }
         case <-safeReceive(tickers["processor"]):
             if config.AppConfig.Drill.Processor != nil && config.AppConfig.Drill.Processor.Enabled {
                 go func() { ds.runDrill("processor") }()
             }
         case <-safeReceive(tickers["address"]):
             if config.AppConfig.Drill.Address != nil && config.AppConfig.Drill.Address.Enabled {
                 go func() { ds.runDrill("address") }()
             }
         case <-safeReceive(tickers["storage"]):
             if config.AppConfig.Drill.Storage != nil && config.AppConfig.Drill.Storage.Enabled {
                 go func() { ds.runDrill("storage") }()
             }
         case <-safeReceive(tickers["disk"]):
             if config.AppConfig.Drill.Disk != nil && config.AppConfig.Drill.Disk.Enabled {
                 go func() { ds.runDrill("disk") }()
             }
         case <-safeReceive(tickers["tape"]):
             if config.AppConfig.Drill.Tape != nil && config.AppConfig.Drill.Tape.Enabled {
                 go func() { ds.runDrill("tape") }()
             }
         case <-safeReceive(tickers["bus"]):
             if config.AppConfig.Drill.Bus != nil && config.AppConfig.Drill.Bus.Enabled {
                 go func() { ds.runDrill("bus") }()
             }
         case <-safeReceive(tickers["network"]):
             if config.AppConfig.Drill.Network != nil && config.AppConfig.Drill.Network.Enabled {
                 go func() { ds.runDrill("network") }()
             }
         case <-safeReceive(tickers["display"]):
             if config.AppConfig.Drill.Display != nil && config.AppConfig.Drill.Display.Enabled {
                 go func() { ds.runDrill("display") }()
             }
         case <-safeReceive(tickers["input"]):
             if config.AppConfig.Drill.Input != nil && config.AppConfig.Drill.Input.Enabled {
                 go func() { ds.runDrill("input") }()
             }
         case <-safeReceive(tickers["printer"]):
             if config.AppConfig.Drill.Printer != nil && config.AppConfig.Drill.Printer.Enabled {
                 go func() { ds.runDrill("printer") }()
             }
         case <-safeReceive(tickers["multimedia"]):
             if config.AppConfig.Drill.Multimedia != nil && config.AppConfig.Drill.Multimedia.Enabled {
                 go func() { ds.runDrill("multimedia") }()
             }
         case <-safeReceive(tickers["communication"]):
             if config.AppConfig.Drill.Communication != nil && config.AppConfig.Drill.Communication.Enabled {
                 go func() { ds.runDrill("communication") }()
             }
         case <-safeReceive(tickers["power"]):
             if config.AppConfig.Drill.Power != nil && config.AppConfig.Drill.Power.Enabled {
                 go func() { ds.runDrill("power") }()
             }
         case <-safeReceive(tickers["volume"]):
             if config.AppConfig.Drill.Volume != nil && config.AppConfig.Drill.Volume.Enabled {
                 go func() { ds.runDrill("volume") }()
             }
         case <-safeReceive(tickers["generic"]):
             if config.AppConfig.Drill.Generic != nil && config.AppConfig.Drill.Generic.Enabled {
                 go func() { ds.runDrill("generic") }()
             }
         }
    }
    select {}
}

func safeReceive(ticker *time.Ticker) <-chan time.Time {
    if ticker != nil {
        return ticker.C
    }
    return nil
}

func (ds *DrillScheduler) runDrill(deviceClass string) {
    jsonData, err := drill.Drill(deviceClass)
    if err != nil {
        fmt.Println("%s Drill Error:", deviceClass, err)
        return
    }
    if deviceClass == "battery" {
        transformer.GetBatteryDetails(jsonData)
    } else {
        transformer.GetDeviceDetails(jsonData)
    }
}

func getDrillConfig(drillConfig *config.DrillConfig, drillHead string) bool {
    if drillConfig == nil {
        return false
    }

    val := reflect.ValueOf(drillConfig).Elem()
    field := val.FieldByName(strings.Title(drillHead))
    if !field.IsValid() || field.IsNil() {
        return false
    }
    return field.Elem().FieldByName("Enabled").Bool()
}
