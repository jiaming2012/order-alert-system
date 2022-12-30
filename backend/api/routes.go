package api

import (
	"fmt"
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

func SetupRoutes() {
	http.HandleFunc("/", renderHomepage)
	http.HandleFunc("/thank-you", renderAsset("template/thank-you.html", "text/html"))
	http.HandleFunc("/login", login)
	http.HandleFunc("/contact_form_style.css", renderContactFormTemplate)
	http.HandleFunc("/assets/thank-you.css", renderAsset("assets/thank-you.css", "text/css"))
	http.HandleFunc("/assets/particles.js", renderAsset("assets/particles.js", "text/javascript"))
	http.HandleFunc("/assets/particles-min-script.js", renderAsset("assets/particles-min-script.js", "text/javascript"))
	http.HandleFunc("/assets/logo.jpg", renderAsset("assets/logo.jpg", "image/jpg"))
	http.HandleFunc("/order", PlaceNewOrder)

	// todo: add auth
	http.HandleFunc("/admin/order", PlaceOrderUpdate)

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		pool := websocket.NewPool()
		go pool.Start()

		serveWs(pool, w, r)
	})
}
