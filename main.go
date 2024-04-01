package main

import (
	"github.com/MarceloJbCosta/gopportunities/config"
	"github.com/MarceloJbCosta/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//initializer configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialzation Error: %v", err)
		return

	}
	//initializrr router
	router.Initializer()
}
