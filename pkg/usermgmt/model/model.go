package model

import (
	// "gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// const (
// 	dbUser     string = "test"
// 	dbPassword string = "ZHOUjian.22"
// 	dbHost     string = "127.0.0.1"
// 	dbPort     int    = 33577
// 	dbName     string = "test"
// )

// var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
// 	dbUser, dbPassword, dbHost, dbPort, dbName)

var models = []interface{}{}

func RegisterModel(model interface{}) {
	models = append(models, model)
}

func GetModels() []interface{} {
	return models
}

func CreateConnection() (*gorm.DB, error) {
	// db, err := gorm.Open(mysql.New(mysql.Config{
	// 	DSN: dsn
	// }), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open("userdb.sqlite3"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
