package subscribers

import (
	"encoding/json"
	"log"
	"weatherdredger/internal/config/structs"

	"github.com/gorilla/websocket"
	"github.com/nats-io/nats.go"
)

func SubscribeTemperature(nc *nats.Conn, ws *websocket.Conn) *nats.Subscription {
	// Subscribe to "weather-temperature"
	Sub, err := nc.Subscribe("weather-temperature", func(msg *nats.Msg) {
		env := structs.Envelope{
			Type: "temperature",
			Data: msg.Data,
		}
		out, err := json.Marshal(env)
		if err != nil {
			log.Println("Error marshaling temperature envelope:", err)
			return
		}
		if writeErr := ws.WriteMessage(websocket.TextMessage, out); writeErr != nil {
			log.Println("Error writing temperature to WebSocket:", writeErr)
		}
	})
	if err != nil {
		log.Println("NATS subscribe to weather-temperature failed:", err)
		return nil
	}
	return Sub
}
