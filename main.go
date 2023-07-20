package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	
	e := echo.New()

 
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Adi's First API code")
	})

	
	e.Start(":8080")
}
