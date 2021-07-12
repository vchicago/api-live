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

package database

import (
	"fmt"
	log2 "log"
	"os"
	"time"

	"github.com/dhawton/log4g"
	kzdvTypes "github.com/vchicago/types/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var MaxAttempts = 10
var DelayBetweenAttempts = time.Minute * 1
var attempt = 1
var log = log4g.Category("db")

func Connect(user string, pass string, hostname string, port string, database string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", user, pass, hostname, port, database)
	newLogger := logger.New(
		log2.New(os.Stdout, "\r\n", log2.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)

	if err != nil {
		log.Error("Error connecting to database: " + err.Error())
		if attempt < MaxAttempts {
			log.Info(fmt.Sprintf("Attempt %d/%d Failed. Waiting %s before trying again...", attempt, MaxAttempts, DelayBetweenAttempts.String()))
			time.Sleep(DelayBetweenAttempts)
			attempt += 1
			Connect(user, pass, hostname, port, database)
			return
		}
		panic("Max attempts occured. Aborting startup.")
	}

	db.AutoMigrate(&kzdvTypes.Flights{})

	DB = db
}
