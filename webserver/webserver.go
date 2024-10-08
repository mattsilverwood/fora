package webserver

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func RunForaServer() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	e := echo.New()
	e.HideBanner = true

	e.GET("/", func(c echo.Context) error {
		log.Info().Str("Path", c.Path()).Msg("Got request")
		return c.String(http.StatusOK, c.Path())
	})

	e.Logger.Fatal(e.Start(":8080"))
}
