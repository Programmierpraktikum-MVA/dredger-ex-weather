package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"
	"weatherdredger/internal/config/structs"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	for {
		reading := structs.TemperatureReading{
			StationID:    "temperature-station-1",
			Timestamp:    time.Now().UTC().Format(time.RFC3339),
			TemperatureC: 15 + rand.Float64()*10,
		}

		data, err := json.Marshal(reading)
		if err != nil {
			log.Println("Error marshaling temperature:", err)
			continue
		}

		err = nc.Publish("weather-temperature", data)
		if err != nil {
			log.Println("Error publishing temperature:", err)
		} else {
			log.Println("Published temperature:", string(data))
		}

		time.Sleep(2 * time.Second)
	}
}
