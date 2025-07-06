package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins — tighten for production
	},
}

// RegisterHandlers sets up Echo routes including WebSocket
func RegisterHandlers(e *echo.Echo) {
	// WebSocket endpoint
	e.GET("/ws", func(c echo.Context) error {
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			log.Println("WebSocket upgrade failed:", err)
			return err
		}
		defer ws.Close()

		// Connect to NATS
		nc, err := nats.Connect(nats.DefaultURL)
		if err != nil {
			log.Println("Failed to connect to NATS:", err)
			return err
		}
		defer nc.Drain()

		// Subscribe to NATS topics
		var subscriptions []*nats.Subscription

		
		subscribetohumidityreadingssub := subscribeToHumidityReadings(nc, ws)
		if subscribetohumidityreadingssub == nil {
			log.Println("Subscription to weather-humidity failed")
			return nil
		}
		subscriptions = append(subscriptions, subscribetohumidityreadingssub)
		
		subscribetotemperaturereadingssub := subscribeToTemperatureReadings(nc, ws)
		if subscribetotemperaturereadingssub == nil {
			log.Println("Subscription to weather-temperature failed")
			return nil
		}
		subscriptions = append(subscriptions, subscribetotemperaturereadingssub)
		

		// Keep WebSocket open until client closes
		for {
			if _, _, err := ws.ReadMessage(); err != nil {
				// Unsubscribe all on WebSocket close
				for _, sub := range subscriptions {
					sub.Unsubscribe()
				}
				break
			}
		}
		return nil
	})

	// Optionally serve static files if needed
	e.Static("/pages", "web/pages")
}