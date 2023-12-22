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
    "bus",
    "memory",
    "processor",
    "bridge",
    "hub",
    "display",
    "multimedia",
    "network",
    "communication",
    "storage",
    "volume",
    "disk",
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
                 ds.runDeviceDrill(dt)
             }(drillHead)
         }
     }

    for {
      select {
        case <-safeReceive(tickers["bus"]):
            if config.AppConfig.Drill.Bus != nil && config.AppConfig.Drill.Bus.Enabled {
                go func() { ds.runDeviceDrill("bus") }()
            }
        case <-safeReceive(tickers["memory"]):
            if config.AppConfig.Drill.Memory != nil && config.AppConfig.Drill.Memory.Enabled {
                go func() { ds.runDeviceDrill("memory") }()
            }
        case <-safeReceive(tickers["processor"]):
            if config.AppConfig.Drill.Processor != nil && config.AppConfig.Drill.Processor.Enabled {
                go func() { ds.runDeviceDrill("processor") }()
            }
        case <-safeReceive(tickers["bridge"]):
            if config.AppConfig.Drill.Bridge != nil && config.AppConfig.Drill.Bridge.Enabled {
                go func() { ds.runDeviceDrill("bridge") }()
            }
        case <-safeReceive(tickers["hub"]):
            if config.AppConfig.Drill.Hub != nil && config.AppConfig.Drill.Hub.Enabled {
                go func() { ds.runDeviceDrill("hub") }()
            }
        case <-safeReceive(tickers["display"]):
            if config.AppConfig.Drill.Display != nil && config.AppConfig.Drill.Display.Enabled {
                go func() { ds.runDeviceDrill("display") }()
            }
        case <-safeReceive(tickers["multimedia"]):
            if config.AppConfig.Drill.Multimedia != nil && config.AppConfig.Drill.Multimedia.Enabled {
                go func() { ds.runDeviceDrill("multimedia") }()
            }
        case <-safeReceive(tickers["network"]):
            if config.AppConfig.Drill.Network != nil && config.AppConfig.Drill.Network.Enabled {
                go func() { ds.runDeviceDrill("network") }()
            }
        case <-safeReceive(tickers["communication"]):
            if config.AppConfig.Drill.Communication != nil && config.AppConfig.Drill.Communication.Enabled {
                go func() { ds.runDeviceDrill("communication") }()
            }
        case <-safeReceive(tickers["storage"]):
            if config.AppConfig.Drill.Storage != nil && config.AppConfig.Drill.Storage.Enabled {
                go func() { ds.runDeviceDrill("storage") }()
            }
        case <-safeReceive(tickers["volume"]):
            if config.AppConfig.Drill.Volume != nil && config.AppConfig.Drill.Volume.Enabled {
                go func() { ds.runDeviceDrill("volume") }()
            }
        case <-safeReceive(tickers["disk"]):
            if config.AppConfig.Drill.Disk != nil && config.AppConfig.Drill.Disk.Enabled {
                go func() { ds.runDeviceDrill("disk") }()
            }
        case <-safeReceive(tickers["generic"]):
            if config.AppConfig.Drill.Generic != nil && config.AppConfig.Drill.Generic.Enabled {
                go func() { ds.runDeviceDrill("generic") }()
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

func (ds *DrillScheduler) runDeviceDrill(deviceClass string) {
    jsonData, err := drill.Drill(deviceClass)
    if err != nil {
        fmt.Println("%s Drill Error:", deviceClass, err)
        return
    }
    transformer.GetDeviceDetails(jsonData)
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
