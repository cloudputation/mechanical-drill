package config

import (
    "fmt"
    "os"

    "github.com/hashicorp/hcl/v2/gohcl"
    "github.com/hashicorp/hcl/v2/hclparse"
    "github.com/hashicorp/nomad/api"
    "github.com/spf13/viper"
)

type Configuration struct {
    LogDir    string `hcl:"log_dir"`
    DataDir   string `hcl:"data_dir"`
    Listen    string `hcl:"listen"`
    NomadHost string `hcl:"nomad_host"`
    Drill     DrillConfig `hcl:"drill,block"`
}

type DrillConfig struct {
    System        *SystemDrillConfig        `hcl:"system,block"`
    Bridge        *BridgeDrillConfig        `hcl:"bridge,block"`
    Battery       *BatteryDrillConfig       `hcl:"battery,block"`
    Memory        *MemoryDrillConfig        `hcl:"memory,block"`
    Processor     *ProcessorDrillConfig     `hcl:"processor,block"`
    Address       *AddressDrillConfig       `hcl:"address,block"`
    Storage       *StorageDrillConfig       `hcl:"storage,block"`
    Disk          *DiskDrillConfig          `hcl:"disk,block"`
    Tape          *TapeDrillConfig          `hcl:"tape,block"`
    Bus           *BusDrillConfig           `hcl:"bus,block"`
    Network       *NetworkDrillConfig       `hcl:"network,block"`
    Display       *DisplayDrillConfig       `hcl:"display,block"`
    Input         *InputDrillConfig         `hcl:"input,block"`
    Printer       *PrinterDrillConfig       `hcl:"printer,block"`
    Multimedia    *MultimediaDrillConfig    `hcl:"multimedia,block"`
    Communication *CommunicationDrillConfig `hcl:"communication,block"`
    Power         *PowerDrillConfig         `hcl:"power,block"`
    Volume        *VolumeDrillConfig        `hcl:"volume,block"`
    Generic       *GenericDrillConfig       `hcl:"generic,block"`
}

type SystemDrillConfig struct {
    Enabled   bool `hcl:"enabled"`
    Frequency int  `hcl:"frequency"`
}

type BatteryDrillConfig struct {
    Enabled   bool `hcl:"enabled"`
    Frequency int  `hcl:"frequency"`
}

type BridgeDrillConfig struct {
    Enabled   bool `hcl:"enabled"`
    Frequency int  `hcl:"frequency"`
}

type MemoryDrillConfig struct {
    Enabled   bool `hcl:"enabled"`
    Frequency int  `hcl:"frequency"`
}

type ProcessorDrillConfig struct {
    Enabled   bool `hcl:"enabled"`
    Frequency int  `hcl:"frequency"`
}

type AddressDrillConfig struct {
    Enabled   bool `hcl:"enabled"`
    Frequency int  `hcl:"frequency"`
}

type StorageDrillConfig struct {
    Enabled   bool `hcl:"enabled"`
    Frequency int  `hcl:"frequency"`
}

type DiskDrillConfig struct {
    Enabled   bool `hcl:"enabled"`
    Frequency int  `hcl:"frequency"`
}

type TapeDrillConfig struct {
    Enabled   bool `hcl:"enabled"`
    Frequency int  `hcl:"frequency"`
}

type BusDrillConfig struct {
    Enabled   bool `hcl:"enabled"`
    Frequency int  `hcl:"frequency"`
}

type NetworkDrillConfig struct {
    Enabled   bool `hcl:"enabled"`
    Frequency int  `hcl:"frequency"`
}

type DisplayDrillConfig struct {
    Enabled   bool `hcl:"enabled"`
    Frequency int  `hcl:"frequency"`
}

type InputDrillConfig struct {
    Enabled   bool `hcl:"enabled"`
    Frequency int  `hcl:"frequency"`
}

type PrinterDrillConfig struct {
    Enabled   bool `hcl:"enabled"`
    Frequency int  `hcl:"frequency"`
}

type MultimediaDrillConfig struct {
    Enabled   bool `hcl:"enabled"`
    Frequency int  `hcl:"frequency"`
}

type CommunicationDrillConfig struct {
    Enabled   bool `hcl:"enabled"`
    Frequency int  `hcl:"frequency"`
}

type PowerDrillConfig struct {
    Enabled   bool `hcl:"enabled"`
    Frequency int  `hcl:"frequency"`
}

type VolumeDrillConfig struct {
    Enabled   bool `hcl:"enabled"`
    Frequency int  `hcl:"frequency"`
}

type GenericDrillConfig struct {
    Enabled   bool `hcl:"enabled"`
    Frequency int  `hcl:"frequency"`
}

var AppConfig Configuration
var NomadClient *api.Client
var ConfigPath string

func LoadConfiguration() error {
    configPath := fmt.Sprintf("/etc/mechanical-drill/config.hcl")
    viper.SetDefault("ConfigPath", configPath)
    viper.BindEnv("ConfigPath", "MD_CONFIG_FILE_PATH")
    ConfigPath = viper.GetString("ConfigPath")

    data, err := os.ReadFile(ConfigPath)
    if err != nil {
        return fmt.Errorf("Failed to read configuration file: %v", err)
    }

    parser := hclparse.NewParser()
    file, diags := parser.ParseHCL(data, ConfigPath)
    if diags.HasErrors() {
        return fmt.Errorf("Failed to parse configuration: %v", diags)
    }

    diags = gohcl.DecodeBody(file.Body, nil, &AppConfig)
    if diags.HasErrors() {
        return fmt.Errorf("Failed to apply configuration: %v", diags)
    }

    err = InitializeNomadClient()
    if err != nil {
        return fmt.Errorf("Error initializing Nomad client: %v", err)
    }


   return nil
}

func InitializeNomadClient() error {
    nomadConfig := api.DefaultConfig()
    nomadConfig.Address = AppConfig.NomadHost

    var err error
    NomadClient, err = api.NewClient(nomadConfig)
    if err != nil {
        return fmt.Errorf("Error creating Nomad client: %v", err)
    }


    return nil
}
