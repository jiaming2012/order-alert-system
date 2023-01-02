package database

import (
	"fmt"
	//"github.com/jiaming2012/order-alert-system/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

const dbHost = "localhost"
const dbUser = "postgres"
const dbPass = "b320saFs"
const dbName = "customer_orders"

var db *gorm.DB
var mutex sync.Mutex

func Setup() error {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", dbHost, dbUser, dbPass, dbName)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrate the schema

	return nil
}

func GetDB() *gorm.DB {
	mutex.Lock()
	return db
}

func ReleaseDB() {
	mutex.Unlock()
}
