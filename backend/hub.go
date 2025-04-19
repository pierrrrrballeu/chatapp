package main

type Hub struct {
	register   chan *Client
	unregister chan *Client
	broadcast  chan *Message
	clientcount  chan int
}

func NewHub() *Hub {
	return &Hub{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *Message),
		clientcount:  make(chan int),
	}
}

func (h *Hub) start() {
	clients := make(map[*Client]bool)

	for {
		select {
		case client := <-h.register:
			clients[client] = true
		case client := <-h.unregister:
			close(client.send)
			delete(clients, client)
		case message := <-h.broadcast:
			for client := range clients {
				if message.sender == client {
					continue
				}

				select {
				case client.send <- message:
				default:
					delete(clients, client)
					close(client.send)
				}
			}
		case h.clientcount <- len(clients):
		}
		
	}
}
