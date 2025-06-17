package structs

// HumidityReading matches the “weatherHumidity” payload in AsyncAPI
type HumidityReading struct {
	StationID       string  `json:"station_id"`
	Timestamp       string  `json:"timestamp"`
	HumidityPercent float64 `json:"humidity_percent"`
}
