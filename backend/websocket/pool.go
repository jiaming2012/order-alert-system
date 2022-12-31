package websocket

import (
	"fmt"
	"github.com/jiaming2012/order-alert-system/backend/models"
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
				fmt.Println("failed to send all orders to client. closing connection ...")
				if wsErr := client.Conn.Close(); wsErr != nil {
					fmt.Println("failed to close the connection")
				}
				continue
			}
			pool.Clients[client] = true
			fmt.Println("a new client joined the pool, len=", len(pool.Clients))

		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("a client left the pool, len=", len(pool.Clients))

		case orders := <-pool.Broadcast:
			fmt.Println("Sending orders to all clients in Pool")
			for cli, _ := range pool.Clients {
				if err := cli.Conn.WriteJSON(orders); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
