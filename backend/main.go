package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jiaming2012/order-alert-system/backend/api"
	"github.com/jiaming2012/order-alert-system/backend/database"
	"github.com/jiaming2012/order-alert-system/backend/models"
	"github.com/jiaming2012/order-alert-system/backend/pubsub"
)

const PORT = 8080

func main() {
	fmt.Println("Order Messenger App v0.01")

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

	fmt.Println("Event bus setup complete!")

	router := gin.Default()

	api.SetupRoutes(router)

	fmt.Printf("listening on :%d\n", PORT)

	router.Run(fmt.Sprintf(":%d", PORT))
	//api.SetupRoutes()

	//http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil)
}
