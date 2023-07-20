package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create an instance of Echo
	e := echo.New()

	// Define a route handler for the root path "/"
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Adi's First API code")
	})

	// Start the server
	e.Start(":8080")
}
