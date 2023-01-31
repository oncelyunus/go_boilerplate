package main

import (
	"fmt"
	"os"

	"github.com/oncelyunus/go_boilerplate/config"
	"github.com/oncelyunus/go_boilerplate/pkg/app"
	"github.com/oncelyunus/go_boilerplate/pkg/utils"
)

func main() {
	configPath := utils.GetConfigPath(os.Getenv("config"))
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		fmt.Errorf("Loading config: %v", err)
	}
	application, err := app.NewApp(cfg)
	if err != nil {
		panic(err)
	}

	err = application.Init()
	if err != nil {
		panic(err)
	}

	// start http server
	application.StartServer()
}
