package main

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"order-service/internal/middleware/httplogger"
	"order-service/internal/router"

	"github.com/joho/godotenv"
	"order-service/internal/config"
)

const (
	appConfigPath = `config/app.env`
)

const (
	errorWhileLoadingAppConfig      = `Error while loading app config`
	errorWhileInitializingAppConfig = `Error while initializing app config`
)

// main Start of the application
func main() {
	//For loading local dev env variables
	_ = godotenv.Load()

	// For loading common env variables
	err := godotenv.Load(appConfigPath)
	exitOnError(err, errorWhileLoadingAppConfig)
	err = config.InitConfig() // For Initializing Struct for env variables
	exitOnError(err, errorWhileInitializingAppConfig)
	log.SetLevel(httplogger.SetLoggerLevelFromConfig(config.GetConfig().LogLevel))
	log.Infof("Logging on log level: %v", log.GetLevel())
	// Starting Rest Server
	router.RunRest()
}

// exitOnError - method which logs the error for exit scenarios
func exitOnError(err error, context string) {
	if err != nil {
		wrappedError := errors.Wrap(err, context)
		log.Fatal(wrappedError)
	}
}
