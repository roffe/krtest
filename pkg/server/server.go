package server

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebsocketServer handler
type WebsocketServer struct {
	clients   map[*websocket.Conn]bool
	broadcast chan Message
	upgrader  websocket.Upgrader
	unhealthy bool
}

// Define our message object
type Message struct {
	Message string `json:"message"`
}

func (self *WebsocketServer) handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-self.broadcast
		// Send it out to every client that is currently connected
		for client := range self.clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(self.clients, client)
			}
		}
	}
}

func (self *WebsocketServer) healthz(w http.ResponseWriter, r *http.Request) {
	if self.unhealthy {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (self *WebsocketServer) handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := self.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()
	// Register our new client
	self.clients[ws] = true
	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(self.clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		self.broadcast <- msg
	}
}

func (self *WebsocketServer) Run(ctx context.Context, quitChan chan bool) error {

	// Configure websocket route
	http.HandleFunc("/ws", self.handleConnections)
	http.HandleFunc("/healthz", self.healthz)
	go self.handleMessages()

	log.Println("http server started on :8000")
	go func() {
		err := http.ListenAndServe(":8000", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
	select {
	case <-ctx.Done():
		return nil
	case <-quitChan:
		self.unhealthy = true
		msg := Message{
			Message: "Server is shutting down in 30 seconds!",
		}
		self.broadcast <- msg
	}
	return nil
}

// New new server
func New() *WebsocketServer {
	var clients = make(map[*websocket.Conn]bool) // connected clients
	var broadcast = make(chan Message)           // broadcast channel
	// Configure the upgrader
	var upgrader = websocket.Upgrader{}
	return &WebsocketServer{
		clients:   clients,
		broadcast: broadcast,
		upgrader:  upgrader,
	}
}
