/*
   ZAU API - Live flights/pilots
   Copyright (C) 2021  Daniel A. Hawton <daniel@hawton.org>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

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
