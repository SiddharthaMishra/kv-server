package main

// Hub maintains the set of active clients and broadcasts messages to them
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound changes to broadcast
	broadcast chan *Pair

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan *Pair),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		// register new socket
		case client := <-h.register:
			h.clients[client] = true
		// unregister a socket
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		// transmit messages to all sockets
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					h.unregister <- client
				}
			}
		}
	}
}
