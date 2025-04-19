package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(c *gin.Context, h *Hub) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		hub:      h,
		conn:     conn,
		send:     make(chan *Message),
		username: "匿名ユーザー" + strconv.Itoa(<-h.clientcount + 1),
	}

	h.register <- client

	go client.readMessage()
	go client.writeMessage()
}

func main() {
	r := gin.Default()
	hub := NewHub()
	go hub.start()

	r.GET("/ws", func(ctx *gin.Context) {
		WsHandler(ctx, hub)
	})

	r.Run(":3000")
}
