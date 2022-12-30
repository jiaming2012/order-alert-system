package models

import (
	"github.com/jiaming2012/order-alert-system/backend/database"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	CreatedAt   time.Time `json:"createdAt"`
	OrderNumber string    `json:"orderNumber"`
	PhoneNumber string    `json:"phoneNumber"`
	Status      string    `json:"status"`
}

func (o *Order) Create() error {
	db := database.GetDB()
	defer database.ReleaseDB()

	tx := db.Create(o)
	return tx.Error
}

func (o *Order) Save() error {
	db := database.GetDB()
	defer database.ReleaseDB()

	tx := db.Save(o)
	return tx.Error
}

func GetOrder(id string) (Order, error) {
	db := database.GetDB()
	defer database.ReleaseDB()

	var order Order

	tx := db.Find(&order, id)
	return order, tx.Error
}

func GetOpenOrders() ([]Order, error) {
	db := database.GetDB()
	defer database.ReleaseDB()

	var orders []Order

	tx := db.Where("status IN ?", []string{"open", "awaiting_pickup"}).Find(&orders)
	return orders, tx.Error
}
