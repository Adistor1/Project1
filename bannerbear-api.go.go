package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tidwall/gjson"
)

const (
	bannerbearAPIKey   = "YOUR_BANNERBEAR_API_KEY"
	bannerbearTemplate = "YOUR_BANNERBEAR_TEMPLATE_ID"
)

type BannerbearRequest struct {
	Text string `json:"text"`
}

type BannerbearResponse struct {
	ImageURL string `json:"image_url"`
}

func generateBannerbearImage(text string) (string, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.bannerbear.com/v2/images/%s/rendered", bannerbearTemplate)
	payload := BannerbearRequest{
		Text: text,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bannerbearAPIKey))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Bannerbear API request failed with status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	imageURL := gjson.GetBytes(body, "image_url").String()
	return imageURL, nil
}

func handleGenerateImage(c echo.Context) error {
	req := new(BannerbearRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	imageURL, err := generateBannerbearImage(req.Text)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := BannerbearResponse{
		ImageURL: imageURL,
	}
	return c.JSON(http.StatusOK, response)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/generate", handleGenerateImage)

	e.Logger.Fatal(e.Start(":8080"))
}
