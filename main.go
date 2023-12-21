package main

import (
		_ "net/http/pprof"
		"net/http"


		"github.com/cloudputation/mechanical-drill/packages/config"
		"github.com/cloudputation/mechanical-drill/packages/engine"
		l "github.com/cloudputation/mechanical-drill/packages/logger"
)


func main() {
	go func() {
			http.ListenAndServe("localhost:6060", nil)
	}()

	err := config.LoadConfiguration()
	if err != nil {
			l.Fatal("Failed to load configuration: %v", err)
	}

	engine.StartEngine()
}
