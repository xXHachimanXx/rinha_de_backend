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
	router.Init()
	config.DatabaseInit()
}
