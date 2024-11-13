package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

var (
	clientID     = os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")
	redirectURL  = os.Getenv("SPOTIFY_REDIRECT_URL")
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	ExpiresIn int `json:"expires_in"`
}

func getSpotifyToken() (string, error) {
	client := resty.New()
	client.SetTimeout(10 * time.Second)

	resp, err := client.R().SetHeader("Content-Type", "application/x-www-form-urlencoded").SetBasicAuth(clientID, clientSecret).SetFormData(map[string]string{
		"grant_type": "client_credentials",
	}).Post("https://accounts.spotify.com/api/token")

	if err != nil {
		return "", fmt.Errorf("Failed to request token: %v", err)
	}

	if resp.IsError() {
		return "", fmt.Errorf("error response from Spotify: %s", resp.String())
	}

	var tokenResp TokenResponse
	if err := json.Unmarshal(resp.Body(), &tokenResp); err != nil {
		return "", fmt.Errorf("failed to parse token response: %v", err)
	}

	return tokenResp.AccessToken, nil
}

func testSpotifyToken() {

}
