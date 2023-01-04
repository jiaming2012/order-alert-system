package websocket

import (
	"github.com/jiaming2012/order-alert-system/backend/models"
	"github.com/sirupsen/logrus"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan []models.Order
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []models.Order),
	}
}

func (pool *Pool) BroadcastAllOrders() error {
	orders, err := models.GetOpenOrders()
	if err != nil {
		return err
	}

	pool.Broadcast <- orders
	return nil
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			if err := SendAllOrders(client); err != nil {
				logrus.Error("failed to send all orders to client. closing connection ...")
				if wsErr := client.Conn.Close(); wsErr != nil {
					logrus.Error("failed to close the connection")
				}
				continue
			}
			pool.Clients[client] = true
			logrus.Info("a new client joined the pool, len=", len(pool.Clients))

		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			logrus.Info("a client left the pool, len=", len(pool.Clients))

		case orders := <-pool.Broadcast:
			logrus.Info("Sending orders to all clients in Pool")
			for cli, _ := range pool.Clients {
				if err := cli.Conn.WriteJSON(orders); err != nil {
					logrus.Error(err)
					return
				}
			}
		}
	}
}
