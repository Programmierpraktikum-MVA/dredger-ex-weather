// Edit this file, as it is a specific handler function for your service
package rest

import (
	"asyncservice/core/log"
	"asyncservice/core/tracing"
	"asyncservice/usecases"

	"net/http"

	"github.com/labstack/echo/v4"
)

// List all weather stations
func ListStations(c echo.Context) error {
	// trace span
	ctx := c.Request().Context()
	ctx, span := tracing.Tracer.Start(ctx, "logMessage")
	defer span.End()

	traceId := span.SpanContext().TraceID().String()
	spanId := span.SpanContext().SpanID().String()
	log.Info().Str("traceId", traceId).Str("spanId", spanId).Str("path", "/").Msg("ListStations")

	// session, err := getSession(c)
	// if err != nil {
	// 	log.Error().Err(err).Msg("ListStations failed")
	// 	return c.NoContent(http.StatusInternalServerError)
	// }

	stations := usecases.ListStations()
	return c.JSON(http.StatusOK, stations)
	// implement your functionality best using a function from a separate file, e.g. usecases/ListStationsDo.go

	// 200 => A list of weather stations
	return c.String(http.StatusNotImplemented, "Temporary handler stub.")
}
