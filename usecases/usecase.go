package usecases

import (
	"errors"
	"sync"
	"time"

	"build/entities"

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
