package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

type Streamer struct {
	gorm.Model
	Name string
	Key  string
}

func addStreamer(streamer Streamer) error {
	return database.Create(&streamer).Error
}

func getStreamerByKey(key string) (streamer Streamer, err error) {
	err = database.First(&streamer, "key = ?", key).Error
	return
}

func getAllStreamers() (streamers []Streamer, err error) {
	err = database.Find(&streamers).Error
	return
}

func init() {
	db, err := gorm.Open(sqlite.Open("/db/wow.db"), &gorm.Config{})
	err = db.AutoMigrate(&Streamer{})
	if err != nil {
		panic(err)
	}
	database = db
}
