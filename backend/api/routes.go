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

//func handlePage(writer http.ResponseWriter, request *http.Request) {
//	if request.Method == "GET" {
//		tmplt, _ = template.ParseFiles("tutorial.html")
//
//		event := News{
//			Headline: "makeuseof.com has everything Tech",
//			Body:     "Visit MUO for anything technology related",
//		}
//
//		err := tmplt.Execute(writer, event)
//
//		if err != nil {
//			return
//		}
//	}
//}

func SetupRoutes() {
	http.HandleFunc("/", renderHomepage)
	http.HandleFunc("/thank-you", renderAsset("template/thank-you.html", "text/html"))
	http.HandleFunc("/login", login)
	http.HandleFunc("/assets/contact_form_style.css", renderAsset("assets/contact_form_style.css", "text/css"))
	http.HandleFunc("/assets/thank-you.css", renderAsset("assets/thank-you.css", "text/css"))
	http.HandleFunc("/400-error.html", renderTemplateWithParams)
	http.HandleFunc("/assets/400-error.css", renderAsset("assets/400-error.css", "text/css"))
	http.HandleFunc("/500-error.html", renderAsset("template/500-error.html", "text/html"))
	http.HandleFunc("/assets/500-error.css", renderAsset("assets/500-error.css", "text/css"))
	http.HandleFunc("/assets/500-error.js", renderAsset("assets/500-error.js", "text/javascript"))
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
