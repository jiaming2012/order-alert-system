package api

import (
	"fmt"
	"github.com/jiaming2012/order-alert-system/backend/models"
	"github.com/jiaming2012/order-alert-system/backend/pubsub"
	"github.com/jiaming2012/order-alert-system/backend/websocket"
	"net/http"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
		return
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func BroadcastOrders(pool *websocket.Pool) func(interface{}) {
	return func(event interface{}) {
		orders, err := models.GetOpenOrders()
		if err != nil {
			fmt.Println("error getting open orders: ", err)
			return
		}

		fmt.Println("broadcasting ...")
		pool.Broadcast <- orders
	}
}

func SetupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()
	if err := pubsub.Subscribe(pubsub.OrderCreated, BroadcastOrders(pool)); err != nil {
		panic(err)
	}
	if err := pubsub.Subscribe(pubsub.OrderUpdated, BroadcastOrders(pool)); err != nil {
		panic(err)
	}

	http.HandleFunc("/", renderHomepage)
	http.HandleFunc("/thank-you.html", renderAsset("templates/thank-you.html", "text/html"))
	http.HandleFunc("/login", login)
	http.HandleFunc("/assets/contact_form_style.css", renderAsset("assets/contact_form_style.css", "text/css"))
	http.HandleFunc("/assets/thank-you.css", renderAsset("assets/thank-you.css", "text/css"))
	http.HandleFunc("/400-error.html", renderTemplateWithParams)
	http.HandleFunc("/assets/400-error.css", renderAsset("assets/400-error.css", "text/css"))
	http.HandleFunc("/500-error.html", renderAsset("templates/500-error.html", "text/html"))
	http.HandleFunc("/assets/500-error.css", renderAsset("assets/500-error.css", "text/css"))
	http.HandleFunc("/assets/500-error.js", renderAsset("assets/500-error.js", "text/javascript"))
	http.HandleFunc("/assets/particles.js", renderAsset("assets/particles.js", "text/javascript"))
	http.HandleFunc("/assets/particles-min-script.js", renderAsset("assets/particles-min-script.js", "text/javascript"))
	http.HandleFunc("/assets/logo.jpg", renderAsset("assets/logo.jpg", "image/jpg"))
	http.HandleFunc("/order", HandlePlaceNewOrder)

	// todo: add auth
	http.HandleFunc("/admin/order", HandlePlaceOrderUpdate)

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}
