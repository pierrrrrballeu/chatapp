package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub      *Hub
	conn     *websocket.Conn
	send     chan *Message
	username string
}

func (c *Client) broadcast(content string, mesasgeType MessageType) {
	c.hub.broadcast <- &Message{
		sender: c,
		data: &ChatMessage{
			Content: content,
			TimeStamp: "17じ37ふん",
			Username: c.username,
			Type: mesasgeType,
		},
	}
}

func (c *Client) readMessage() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
			
	for {
		var msg ChatMessage
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("Received: %s", msg.Content)

		c.broadcast(msg.Content, TextMessage)
	}
}

func (c *Client) writeMessage() {
	for msg := range c.send {
		if err := c.conn.WriteJSON(msg.data); err != nil {
			log.Println(err)
			return
		}
	}
}
