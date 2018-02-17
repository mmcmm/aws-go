package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func main() {

	db, err = gorm.Open("sqlite3", "./gorm.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.LogMode(true)
	db.AutoMigrate(&user{}, &comment{}, &photo{}, &follower{})

	r := registerRoutes()

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	log.Info("Listening on :", port)
	r.Run(":" + port)
}
