package server

import (
	"fmt"
	"webEngineering/pkg/api"
	"webEngineering/pkg/config"
)

func Start() error {
	var cfg config.Config
	cfg.Setup()

	apiKeeper, err := api.New(cfg)
	if err != nil {
		return err
	}

	apiKeeper.Fill()
	fmt.Println("Filled and ready to serve", cfg)
	return apiKeeper.Serve()
}
