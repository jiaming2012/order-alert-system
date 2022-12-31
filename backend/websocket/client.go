package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/jiaming2012/order-alert-system/backend/models"
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

		fmt.Printf("Message Received: %+v\n", message)
	}
}
