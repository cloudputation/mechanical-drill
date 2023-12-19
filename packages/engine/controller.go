package engine

import (
  	"fmt"
  	"time"

    "github.com/cloudputation/mechanical-drill/packages/config"
    "github.com/cloudputation/mechanical-drill/packages/engine/drill"
)


func (ds *DrillScheduler) RunController() {
    switch {
    case config.AppConfig.Drill.PCI.Enabled:
        jsonData, err := drill.DrillPCI()
        if err != nil {
            fmt.Println("Error:", err)
            time.Sleep(time.Duration(ds.PCI) * time.Second)
        }

        PrintDeviceDetails(jsonData)
    default:
        fmt.Println("PCI Drill is disabled in the configuration.")
    }

    time.Sleep(time.Duration(ds.PCI) * time.Second)
}
