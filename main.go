package main

import (
	"event-bus-demo/infrastructure/configuration"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"os"
)

func main() {
	arguments := configuration.ParseInputArguments()
	config, err := loadConfiguration(arguments)
	if err != nil {
		log.Fatalf("failed loading configuration due to %s", err.Error())
	}
	deps, err := initializeDependencies(config)
	if err != nil {
		log.Fatalf("failed while initializing dependencies due to %s", err.Error())
	}
	router := initializeRoutes(arguments.ActiveConfigurationProfiles, deps.RequiredControllers)
	deps.EventBus.Run()
	_ = os.Setenv("PORT", fmt.Sprintf("%d", *config.Gin.Port))
	_ = router.Run()
}

func loadConfiguration(arguments configuration.Arguments) (configuration.ApplicationConfiguration, error) {
	validate := validator.New()
	yamlParser := configuration.NewYamlParser(validate)
	return yamlParser.ReadConfiguration(arguments)
}
