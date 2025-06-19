# Beispielprojekt für AsyncAPI - Microservice zur Wetterabfrage

Dieses Beispielprojekt soll einen Microservice implementieren, der Wetterdaten simuliert und publiziert.

## Kurzanleitung zum Testen:

1. Docker installieren
2. `docker run -d -p 4222:4222 nats`
3. `go run cmd/server/main.go`
4. In zweitem Terminal: `go run cmd/publisher/temperature.go`
5. In dritten Terminal: `go run cmd/publisher/humidity.go`
