package subscribers

import (
	"encoding/json"
	"log"
	"weatherdredger/internal/config/structs"

	"github.com/gorilla/websocket"
	"github.com/nats-io/nats.go"
)

func SubscribeHumidity(nc *nats.Conn, ws *websocket.Conn) *nats.Subscription {

	// Subscribe to "weather-humidity"
	Sub, err := nc.Subscribe("weather-humidity", func(msg *nats.Msg) {
		env := structs.Envelope{
			Type: "humidity",
			Data: msg.Data,
		}
		out, err := json.Marshal(env)
		if err != nil {
			log.Println("Error marshaling humidity envelope:", err)
			return
		}
		if writeErr := ws.WriteMessage(websocket.TextMessage, out); writeErr != nil {
			log.Println("Error writing humidity to WebSocket:", writeErr)
		}
	})
	if err != nil {
		log.Println("NATS subscribe to weather-humidity failed:", err)
		return nil
	}

	return Sub
}
