# Beispielprojekt für AsyncAPI - Microservice zur Wetterabfrage

Dieses Beispielprojekt soll einen Microservice implementieren, der Wetterdaten simuliert und publiziert. 
Features:
# OpenAPI
_Zweck: Integration von sowohl OpenAPI als auch AsyncAPI_
- Erstelle Wetterstationen
- Betrachte Wetterstationen
- Erstelle Updates oder Infos zu Wetterstationen

# AsyncAPI
- Betrachte "Live-Sensordaten" von zwei Wetterstationen
- _TODO: Mehr als einen Nachrichtentyp pro Channel_
- _TODO: Nicht mehr nur NATS_

## Kurzanleitung zum Testen:

1. Docker installieren
2. `docker run -d -p 4222:4222 nats`
3. `go run .` im Projektordner
4. GUI unter `http://localhost:8080`

**NEU:** Die Sensoren werden automatisch gestartet

# Kurzanleitung - Beispielprojekt selbst generieren

1. Specfiles einfügen
2. Dredger laufen lassen mit `go run main.go generate examples/simple/weather-example.json -o ./build-wetter -n async-service -f -D` --Ergänzen--
3. Businesslogik einfügen:
- `publishers/weather-humidity.go`

```go
func publishHumidityReading() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	for {
		reading := entities.HumidityReading{
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
```

- `publishers/weather-temperature.go`

```go
func publishTemperatureReading(){
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()
	time.Sleep(time.Second)
	for {
		reading := entities.TemperatureReading{
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
```

- `mainSvc.go` in `mainSvc()`;

```go
go publishers.Humidity()
go publishers.Temperature()
```