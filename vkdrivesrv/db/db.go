package db

import (
	_ "log"
	_ "time"
	_"fmt"

	_ "gopkg.in/gormigrate.v1"

	"../config"
	//_"../model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dataBase *gorm.DB

func ConnectToDB() *gorm.DB {
	return connect()
}

func DisconnectFromDB() {
	disconnect()
}

func connect() *gorm.DB {
	if dataBase != nil {
		return dataBase
	}
	options := "host=" + config.Config.Database.Host + " user=" + config.Config.Database.User + " dbname=" + config.Config.Database.Dbname + " sslmode=disable password=" + config.Config.Database.Password
	db, err := gorm.Open("postgres", options)
	if err != nil {
		panic("failed to connect database")
	}
	dataBase = db
	return dataBase
}

func disconnect() (error) {
	return dataBase.Close()
}


//CreateTable can create new table from struct interface
func createTable(model interface{}) error {
	db := connect()

	db.CreateTable(model)
	return nil
}