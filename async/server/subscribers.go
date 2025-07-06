package server

import (
	"asyncservice/entities"
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	"github.com/nats-io/nats.go"
)

func subscribeToHumidityReadings(nc *nats.Conn, ws *websocket.Conn) *nats.Subscription {

	// Subscribe to "weather-humidity"
	Sub, err := nc.Subscribe("weather-humidity", func(msg *nats.Msg) {
		env := entities.Envelope{
			Type: "weather-humidity",
			Data: msg.Data,
		}
		out, err := json.Marshal(env)
		if err != nil {
			log.Println("Error marshaling weatherHumidity envelope:", err)
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

func subscribeToTemperatureReadings(nc *nats.Conn, ws *websocket.Conn) *nats.Subscription {

	// Subscribe to "weather-temperature"
	Sub, err := nc.Subscribe("weather-temperature", func(msg *nats.Msg) {
		env := entities.Envelope{
			Type: "weather-temperature",
			Data: msg.Data,
		}
		out, err := json.Marshal(env)
		if err != nil {
			log.Println("Error marshaling weatherTemperature envelope:", err)
			return
		}
		if writeErr := ws.WriteMessage(websocket.TextMessage, out); writeErr != nil {
			log.Println("Error writing humidity to WebSocket:", writeErr)
		}
	})
	if err != nil {
		log.Println("NATS subscribe to weather-temperature failed:", err)
		return nil
	}

	return Sub
}
