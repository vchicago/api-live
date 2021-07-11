package main

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vchicago/api-live/database"

	"github.com/dhawton/log4g"
)

var log = log4g.Category("main")

func main() {
	intro := figure.NewFigure("ZAU Overflight", "", false).Slicify()
	for i := 0; i < len(intro); i++ {
		log.Info(intro[i])
	}

	log.Info("Checking for .env, loading if exists")
	if _, err := os.Stat(".env"); err == nil {
		log.Info("Found, loading")
		err := godotenv.Load()
		if err != nil {
			log.Error("Error loading .env file: " + err.Error())
		}
	}

	appenv := Getenv("APP_ENV", "dev")
	log.Debug(fmt.Sprintf("APPENV=%s", appenv))

	if appenv == "production" {
		log4g.SetLogLevel(log4g.INFO)
		log.Info("Setting gin to Release Mode")
		gin.SetMode(gin.ReleaseMode)
	} else {
		log4g.SetLogLevel(log4g.DEBUG)
	}

	log.Info("Connecting to database and handling migrations")
	database.Connect(Getenv("DB_USERNAME", "root"), Getenv("DB_PASSWORD", "secret"), Getenv("DB_HOSTNAME", "localhost"), Getenv("DB_PORT", "3306"), Getenv("DB_DATABASE", "zau"))

	log.Info("Configuring gin webserver")
	server := NewServer(appenv)

	log.Info("Done, running web server")
	server.engine.Run(fmt.Sprintf(":%s", Getenv("PORT", "3000")))
}
