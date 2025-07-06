package publishers

import (
	"asyncservice/entities"
	"encoding/json"
	"log" //noch an core logger anpassen
	"math/rand"
	"time"

	"github.com/nats-io/nats.go"
)

func PublishTemperatureReading() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()
	time.Sleep(time.Second)
	for {
		reading := entities.WeatherTemperature{
			StationId:    "temperature-station-1",
			Timestamp:    time.Now().UTC().Format(time.RFC3339),
			TemperatureC: 15 + rand.Float64()*10,
		}

		data, err := json.Marshal(reading)
		if err != nil {
			log.Println("Error marshaling temperature:", err)
			continue
		}

		err = nc.Publish("WeatherTemperature", data)
		if err != nil {
			log.Println("Error publishing temperature:", err)
		} else {
			log.Println("Published temperature:", string(data))
		}

		time.Sleep(2 * time.Second)
	}
}
