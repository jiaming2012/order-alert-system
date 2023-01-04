package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jiaming2012/order-alert-system/backend/api"
	"github.com/jiaming2012/order-alert-system/backend/database"
	"github.com/jiaming2012/order-alert-system/backend/models"
	"github.com/jiaming2012/order-alert-system/backend/pubsub"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.Info("Order Messenger App v0.01")

	logrus.Info("Setting up database ...")
	if err := database.Setup(); err != nil {
		logrus.Errorf("failed to setup database: %v", err)
		return
	}
	db := database.GetDB()
	db.AutoMigrate(&models.Order{})
	database.ReleaseDB()
	logrus.Info("Db setup complete!")

	logrus.Info("Setting up event bus ...")
	if err := pubsub.Setup(); err != nil {
		logrus.Errorf("failed to setup event bus: %v", err)
		return
	}

	logrus.Info("Event bus setup complete!")

	router := gin.Default()

	api.SetupRoutes(router)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	logrus.Infof("listening on :%s", port)

	router.Run(fmt.Sprintf(":%s", port))
}
