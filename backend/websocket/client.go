package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/jiaming2012/order-alert-system/backend/models"
	"github.com/sirupsen/logrus"
	"log"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func BroadcastOrders(pool *Pool) func(interface{}) {
	return func(event interface{}) {
		orders, err := models.GetOpenOrders()
		if err != nil {
			logrus.Error("error getting open orders: ", err)
			return
		}

		pool.Broadcast <- orders
	}
}

func SendAllOrders(c *Client) error {
	orders, err := models.GetOpenOrders()
	if err != nil {
		return err
	}
	c.Conn.WriteJSON(orders)
	return nil
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := string(p)

		logrus.Debugf("Message Received: %+v\n", message)
	}
}
