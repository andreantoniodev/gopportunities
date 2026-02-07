package main

import (
	"github.com/andreantoniodev/gopportunities.git/config"
	"github.com/andreantoniodev/gopportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Failed to initialize configuration: %v", err)
		return
	}

	router.Initialize()
}
