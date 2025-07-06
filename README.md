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
func PublishHumidityReading() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	for {
		reading := entities.WeatherHumidity{
			StationId:       "humidity-station-1",
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
func PublishTemperatureReading(){
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

- `usecases/usecase.go`

```go
package usecases

import (
	"errors"
	"sync"
	"time"

	"asyncservice/entities"

	"github.com/google/uuid"
)

var (
	mu       sync.RWMutex
	stations = map[string]*entities.WeatherStation{}
)

// Create a new station
func CreateStation(input entities.NewWeatherStation) (*entities.WeatherStation, error) {
	mu.Lock()
	defer mu.Unlock()

	id := uuid.New().String()

	newStation := &entities.WeatherStation{
		Id:          id,
		Name:        input.Name,
		Type:        input.Type,
		LastReading: entities.WeatherUpdate{},
	}

	stations[id] = newStation
	return newStation, nil
}

// List all stations
func ListStations() []*entities.WeatherStation {
	mu.RLock()
	defer mu.RUnlock()

	result := []*entities.WeatherStation{}
	for _, s := range stations {
		result = append(result, s)
	}
	return result
}

// Get a station by ID
func GetStationByID(id string) (*entities.WeatherStation, error) {
	mu.RLock()
	defer mu.RUnlock()

	if station, ok := stations[id]; ok {
		return station, nil
	}
	return nil, errors.New("station not found")
}

// Update a station reading
func UpdateStationReading(id string, update entities.WeatherUpdate) error {
	mu.Lock()
	defer mu.Unlock()

	station, ok := stations[id]
	if !ok {
		return errors.New("station not found")
	}

	update.Timestamp = time.Now().UTC().Format(time.RFC3339)
	station.LastReading = update
	return nil
}
```
- `mainSvc.go` import;

```go
import (
	"asyncservice/async/publishers"
)
```

- `mainSvc.go` in `mainSvc()`;

```go
go publishers.PublishHumidityReading()
go publishers.PublishTemperatureReading()
```

- `rest/listStations.go`;

```go
stations := usecases.ListStations()
return c.JSON(http.StatusOK, stations)
```

- `rest/updateStationReadings.go`;

```go
err = usecases.UpdateStationReading(stationId, content)
if err != nil {
	return c.String(http.StatusNotFound, err.Error())
}
return c.NoContent(http.StatusOK)
```

- `rest/getStationById.go`;

```go
station, err := usecases.GetStationByID(stationId)
if err != nil {
	return c.String(http.StatusNotFound, err.Error())
}
return c.JSON(http.StatusOK, station)
```

