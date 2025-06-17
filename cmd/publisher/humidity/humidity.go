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
		reading := structs.HumidityReading{
			StationID:       "humidity-station-1",
			Timestamp:       time.Now().UTC().Format(time.RFC3339),
			HumidityPercent: 40 + rand.Float64()*20,
		}

		data, err := json.Marshal(reading)
		if err != nil {
			log.Println("Error marshaling humidity:", err)
			continue
		}

		err = nc.Publish("weather-humidity", data)
		if err != nil {
			log.Println("Error publishing humidity:", err)
		} else {
			log.Println("Published humidity:", string(data))
		}

		time.Sleep(2 * time.Second)
	}
}
