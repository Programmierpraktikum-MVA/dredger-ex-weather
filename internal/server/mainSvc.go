package server

import (
	"log"
	"net/http"

	"weatherdredger/internal/server/subscribers"

	"github.com/gorilla/websocket"
	"github.com/nats-io/nats.go"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all origins — restrict in prod!
	},
}

// RegisterHandlers sets up all HTTP routes
func RegisterHandlers() {
	// Static files
	http.Handle("/", http.FileServer(http.Dir("web/pages")))

	// WebSocket endpoint
	http.HandleFunc("/ws", handleWebSocket)
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer ws.Close()

	// Connect to NATS per connection (or use shared pool)
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Println("Failed to connect to NATS:", err)
		return
	}
	defer nc.Drain()

	// Subscribe to relevant channels
	humSub := subscribers.SubscribeHumidity(nc, ws)
	tempSub := subscribers.SubscribeTemperature(nc, ws)
	if humSub == nil || tempSub == nil {
		log.Println("Subscription failed")
		return
	}

	// Keep WebSocket open
	for {
		if _, _, err := ws.ReadMessage(); err != nil {
			humSub.Unsubscribe()
			tempSub.Unsubscribe()
			break
		}
	}
}
