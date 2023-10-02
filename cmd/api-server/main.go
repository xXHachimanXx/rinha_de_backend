package main

import (
	"github.com/xXHachimanXx/rinha_de_backend/config"
	"github.com/xXHachimanXx/rinha_de_backend/internal/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	router.Init()

}
