// Edit this file, as it is a specific for your service
package main

import (
	"build/async/publishers"
)

func mainSvc() {
	// Add your own evaluation of the flags for additional commands
	// Add your own service running concurrent with the echo web service
	go publishers.Humidity()
	go publishers.Temperature()
}
