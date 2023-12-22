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
    Network     *NetworkDrillConfig     `hcl:"network,block"`
    Storage     *StorageDrillConfig     `hcl:"storage,block"`
    Bus         *BusDrillConfig         `hcl:"bus,block"`
    Memory      *MemoryDrillConfig      `hcl:"memory,block"`
    Processor   *ProcessorDrillConfig   `hcl:"cpu,block"`
    Bridge      *BridgeDrillConfig      `hcl:"bridge,block"`
    Hub         *HubDrillConfig         `hcl:"hub,block"`
    Display     *DisplayDrillConfig     `hcl:"display,block"`
    Multimedia  *MultimediaDrillConfig  `hcl:"multimedia,block"`
    Communication *CommunicationDrillConfig `hcl:"communication,block"`
    Volume      *VolumeDrillConfig      `hcl:"volume,block"`
    Disk        *DiskDrillConfig        `hcl:"disk,block"`
    Generic     *GenericDrillConfig     `hcl:"generic,block"`
}

type NetworkDrillConfig struct {
    Enabled   bool  `hcl:"enabled"`
    Frequency int   `hcl:"frequency"`
}

type StorageDrillConfig struct {
    Enabled   bool  `hcl:"enabled"`
    Frequency int   `hcl:"frequency"`
}

type BusDrillConfig struct {
    Enabled   bool  `hcl:"enabled"`
    Frequency int   `hcl:"frequency"`
}

type MemoryDrillConfig struct {
    Enabled   bool  `hcl:"enabled"`
    Frequency int   `hcl:"frequency"`
}

type ProcessorDrillConfig struct {
    Enabled   bool  `hcl:"enabled"`
    Frequency int   `hcl:"frequency"`
}

type BridgeDrillConfig struct {
    Enabled   bool  `hcl:"enabled"`
    Frequency int   `hcl:"frequency"`
}

type HubDrillConfig struct {
    Enabled   bool  `hcl:"enabled"`
    Frequency int   `hcl:"frequency"`
}

type DisplayDrillConfig struct {
    Enabled   bool  `hcl:"enabled"`
    Frequency int   `hcl:"frequency"`
}

type MultimediaDrillConfig struct {
    Enabled   bool  `hcl:"enabled"`
    Frequency int   `hcl:"frequency"`
}

type CommunicationDrillConfig struct {
    Enabled   bool  `hcl:"enabled"`
    Frequency int   `hcl:"frequency"`
}

type VolumeDrillConfig struct {
    Enabled   bool  `hcl:"enabled"`
    Frequency int   `hcl:"frequency"`
}

type DiskDrillConfig struct {
    Enabled   bool  `hcl:"enabled"`
    Frequency int   `hcl:"frequency"`
}

type GenericDrillConfig struct {
    Enabled   bool  `hcl:"enabled"`
    Frequency int   `hcl:"frequency"`
}


var AppConfig Configuration
var NomadClient *api.Client
var ConfigPath string
var RootDir string


func LoadConfiguration() error {
    RootDir, err := os.Getwd()
    if err != nil {
        return fmt.Errorf("Failed to get service root directory: %v", err)
    }

    configPath := fmt.Sprintf("%s/.release/defaults/config.hcl", RootDir)
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
