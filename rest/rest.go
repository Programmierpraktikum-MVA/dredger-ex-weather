// Don't edit this file, as it is generated by dredger
package rest

import (
	// "asyncservice/rest/middleware"
	"asyncservice/core"
	"asyncservice/core/log"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

var sessionStore *sessions.CookieStore

func init() {
	sessionStore = sessions.NewCookieStore([]byte(core.AppConfig.SessionKey))
}

func getSession(c echo.Context) (*sessions.Session, error) {
	session, err := sessionStore.Get(c.Request(), "session")
	if err != nil || session.IsNew {
		id := uuid.Must(uuid.NewV7()).String()
		log.Info().Str("id", id).Msg("New session")
		session.Values["id"] = id
		//
		// set your default session state
		//
		err = session.Save(c.Request(), c.Response())
		if err != nil {
			log.Error().Err(err).Msg("writing session store failed")
			return nil, err
		}
	}

	return session, nil
}

func NewHandler(e *echo.Echo) {
	g := e.Group("")

	// TODO: Make a validation with openAPI sense
	// 	spec, err := middleware.ParseOpenAPISpecFile("public/weather-api.yaml", e.AcquireContext())
	// 	if err != nil {
	// 		e.Logger.Fatal(err)
	// 	}
	// 	g.Use(middleware.Validation(spec, e))
	//     e.HTTPErrorHandler = func(err error, c echo.Context) {
	//         switch err := err.(type) {
	//         // in case ValidationError just send the message
	//         case middleware.ValidationError:
	//             c.String(err.Status, err.Message)
	//         // in all other cases use default behavior
	//         default:
	//             e.DefaultHTTPErrorHandler(err, c)
	//         }
	//     }

	// Operations for: "/stations"
	g.GET("/stations", ListStations)
	// Operations for: "/stations"
	g.POST("/stations", CreateStation)

	// Operations for: "/stations/:stationId"
	g.GET("/stations/:stationId", GetStationById)
	// Operations for: "/stations/:stationId"
	g.PUT("/stations/:stationId", UpdateStationReading)

	// Call handler extensions
	newHandlerExt(e)
}
