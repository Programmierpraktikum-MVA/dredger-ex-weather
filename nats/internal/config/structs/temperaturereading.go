package structs

// TemperatureReading matches the “weatherTemperature” payload in AsyncAPI
type TemperatureReading struct {
	StationID    string  `json:"station_id"`
	Timestamp    string  `json:"timestamp"`
	TemperatureC float64 `json:"temperature_c"`
}
