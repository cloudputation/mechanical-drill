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
    PCI PCIDrillConfig `hcl:"pci,block"`
}

type PCIDrillConfig struct {
    Enabled   bool  `hcl:"enabled"`
    Frequency int   `hcl:"frequency"`
}

var AppConfig Configuration
var NomadClient *api.Client // Nomad client
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
