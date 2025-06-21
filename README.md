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

**NEU:** Die Sensoren werden automatisch gestartet
