// Edit this file, as it is a specific handler function for your service
package rest

import (
	"asyncservice/core/log"
	"asyncservice/core/tracing"
	"asyncservice/usecases"

	"net/http"

	"github.com/labstack/echo/v4"
)

// Get a specific weather station by ID
func GetStationById(c echo.Context) error {
	// trace span
	ctx := c.Request().Context()
	ctx, span := tracing.Tracer.Start(ctx, "logMessage")
	defer span.End()

	traceId := span.SpanContext().TraceID().String()
	spanId := span.SpanContext().SpanID().String()
	log.Info().Str("traceId", traceId).Str("spanId", spanId).Str("path", "/").Msg("GetStationById")

	// session, err := getSession(c)
	// if err != nil {
	// 	log.Error().Err(err).Msg("GetStationById failed")
	// 	return c.NoContent(http.StatusInternalServerError)
	// }

	stationId := c.Param("stationId")

	station, err := usecases.GetStationByID(stationId)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, station)

	// implement your functionality best using a function from a separate file, e.g. usecases/GetStationByIdDo.go

	// 404 => Station not found
	// 200 => Details of a weather station
	return c.String(http.StatusNotImplemented, "Temporary handler stub.")
}
