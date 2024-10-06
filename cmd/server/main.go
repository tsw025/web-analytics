package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	// We are setting the timezone to UTC, so that all the time values are stored in UTC
	time.Local = time.UTC

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
