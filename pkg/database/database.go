package database

import (
	"fmt"
	"sync"

	"github.com/maoaeri/openapi/pkg/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type singleton struct {
	DB *gorm.DB
}

var instance *singleton
var once sync.Once

func GetDB() *gorm.DB {

	var (
		DB_HOST     = helper.GetEnvVar("DB_HOST")
		DB_PORT     = helper.GetEnvVar("DB_PORT")
		DB_USER     = helper.GetEnvVar("DB_USER")
		DB_PASSWORD = helper.GetEnvVar("DB_PASSWORD")
		DB_NAME     = helper.GetEnvVar("DB_NAME")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqldb, err := connection.DB()
	if err != nil {
		panic(err)
	}

	if err = sqldb.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")
	return connection
}

func GetDBInstance() *singleton {
	connection := GetDB()
	once.Do(func() {
		instance = &singleton{DB: connection}
	})
	return instance
}

func CloseDB(connection *gorm.DB) {
	sqldb, err := connection.DB()
	if err != nil {
		panic(err)
	}
	sqldb.Close()
}
