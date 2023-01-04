package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jiaming2012/order-alert-system/backend/constants"
	"github.com/jiaming2012/order-alert-system/backend/pubsub"
	"github.com/jiaming2012/order-alert-system/backend/websocket"
	"net/http"
)

func serveWs(pool *websocket.Pool, ctx *gin.Context) {
	conn, err := websocket.Upgrade(ctx.Writer, ctx.Request)
	if err != nil {
		sendBadServerErrResponse(err, ctx)
		return
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func SetupRoutes(router *gin.Engine) {
	pool := websocket.NewPool()
	go pool.Start()
	if err := pubsub.Subscribe(pubsub.OrderCreated, websocket.BroadcastOrders(pool)); err != nil {
		panic(err)
	}
	if err := pubsub.Subscribe(pubsub.OrderUpdated, websocket.BroadcastOrders(pool)); err != nil {
		panic(err)
	}

	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	router.GET("/", getHomepage)
	router.GET("/favicon.ico", func(ctx *gin.Context) {
		ctx.File("assets/logo.jpg")
	})
	router.GET("/400-error.html", renderTemplateWithParams)
	router.GET("/thank-you.html", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "thank-you.html", gin.H{})
	})
	router.POST("/", postHomepageForm)
	router.POST("/order", handlePlaceNewOrder)
	router.GET("/orders", func(ctx *gin.Context) {
		serveWs(pool, ctx)
	})

	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		constants.BasicAuthUser: constants.BasicAuthPass,
	}))
	authorized.Use(CORSMiddleware())
	authorized.Static("/", "web")
	// todo: add auth - cannot hide basicAuth creds in a react app
	router.POST("/admin/order", handlePlaceOrderUpdate)
}
