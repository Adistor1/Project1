package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/parnurzeal/gorequest"
)

func main() {
	e := echo.New()

	e.POST("/generate-image", generateImageHandler)

	e.Start(":8080")
}

func generateImageHandler(c echo.Context) error {

	bannerbearAPIKey := "bb_pr_b4b7558b8d41f21b6199c56a3e7bf2"

	input := struct {
		Text string `json:"text"`
	}{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input format")
	}

	templateID := "kY4Qv7D8VJjjZB0qmP"

	parameters := map[string]interface{}{
		"text": input.Text,
	}

	url := fmt.Sprintf("https://api.bannerbear.com/v2/images/%s", templateID)
	request := gorequest.New()
	_, body, errs := request.
		Post(url).
		Set("Authorization", fmt.Sprintf("Bearer %s", bannerbearAPIKey)).
		Send(parameters).
		End()

	if len(errs) > 0 {
		return c.JSON(http.StatusInternalServerError, "Error generating image")
	}

	return c.Blob(http.StatusOK, "image/png", []byte(body))
}
