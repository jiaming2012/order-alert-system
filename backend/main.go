package main

import (
	"fmt"
	"github.com/jiaming2012/order-alert-system/backend/api"
	"github.com/jiaming2012/order-alert-system/backend/database"
	"github.com/jiaming2012/order-alert-system/backend/models"
	"github.com/jiaming2012/order-alert-system/backend/pubsub"
	"github.com/jiaming2012/order-alert-system/backend/sms"
	"net/http"
	"time"
)

const PORT = 8080

func main() {
	fmt.Println("Order Alert System App v0.01")

	fmt.Println("Setting up database ...")
	if err := database.Setup(); err != nil {
		fmt.Printf("failed to setup database: %v", err)
		return
	}
	db := database.GetDB()
	db.AutoMigrate(&models.Order{})
	database.ReleaseDB()
	fmt.Println("Db setup complete!")

	fmt.Println("Setting up event bus ...")
	if err := pubsub.Setup(); err != nil {
		fmt.Printf("failed to setup event bus: %v", err)
		return
	}
	if err := pubsub.Subscribe(pubsub.NewOrderCreated, sms.Send); err != nil {
		fmt.Printf("failed to subscribe to NewOrderCreated event: %v", err)
		return
	}
	//pubsub.Publish(pubsub.NewOrderCreated, pubsub.NewOrderCreatedEvent{
	//	Data: "some data",
	//})
	fmt.Println("Event bus setup complete!")

	test := models.Order{
		OrderNumber: "2",
		PhoneNumber: "856-503-8872",
		CreatedAt:   time.Now(),
		Status:      "open",
	}

	if err := test.Create(); err != nil {
		fmt.Println(err)
		return
	}

	orders, err := models.GetOpenOrders()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, order := range orders {
		fmt.Println(order.OrderNumber, order.PhoneNumber, order.Status, order.ID)
	}

	//if err = sms.SendSMS("+18582227085", "Sent from your Twilio trial account - You rock Natasha! I am working on the order notification system, while waiting for my mom to finish packing. We are going to head back to DC tonight. We have to stop by my Aunts house in Phila first, so it's going to be a late night :("); err != nil {
	//	fmt.Println(err)
	//	return
	//}

	fmt.Println("Setting up http api and websockets ...")
	api.SetupRoutes()
	fmt.Println("Websocket setup complete!")
	http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil)
}
