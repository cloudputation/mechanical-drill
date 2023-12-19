package main

import (
		"github.com/cloudputation/mechanical-drill/packages/config"
		"github.com/cloudputation/mechanical-drill/packages/engine"
		l "github.com/cloudputation/mechanical-drill/packages/logger"
)


func main() {
	err := config.LoadConfiguration()
	if err != nil {
			l.Fatal("Failed to load configuration: %v", err)
	}

	engine.StartEngine()
}
