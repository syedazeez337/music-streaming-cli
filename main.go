package main

import (
    "encoding/base64"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/go-resty/resty/v2"
    "github.com/joho/godotenv"
)

var (
    clientID     = os.Getenv("SPOTIFY_CLIENT_ID")
    clientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")
)

type TokenResponse struct {
    AccessToken string `json:"access_token"`
    TokenType   string `json:"token_type"`
    ExpiresIn   int    `json:"expires_in"`
}

func init() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }
}

// Function to get Spotify access token
func getSpotifyAccessToken() (string, error) {
    client := resty.New()
    client.SetTimeout(10 * time.Second)

    // Encode clientID and clientSecret in Base64
    credentials := fmt.Sprintf("%s:%s", clientID, clientSecret)
    encodedCredentials := base64.StdEncoding.EncodeToString([]byte(credentials))

    // Set form data and headers
    resp, err := client.R().
        SetHeader("Authorization", "Basic "+encodedCredentials).
        SetHeader("Content-Type", "application/x-www-form-urlencoded").
        SetFormData(map[string]string{
            "grant_type": "client_credentials",
        }).
        Post("https://accounts.spotify.com/api/token")

    if err != nil {
        return "", fmt.Errorf("failed to request token: %v", err)
    }

    if resp.IsError() {
        return "", fmt.Errorf("error response from Spotify: %s", resp.String())
    }

    // Parse response
    var tokenResp TokenResponse
    if err := json.Unmarshal(resp.Body(), &tokenResp); err != nil {
        return "", fmt.Errorf("failed to parse token response: %v", err)
    }

    return tokenResp.AccessToken, nil
}

func main() {
    accessToken, err := getSpotifyAccessToken()
    if err != nil {
        log.Fatalf("Failed to get access token: %v", err)
    }
    fmt.Println("Access token received:", accessToken)
}



/*
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
	} else {
		log.Println(".env file loaded successfully")
	}
}

var (
	clientID     = os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")
	redirectURL  = os.Getenv("SPOTIFY_REDIRECT_URL")
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func getSpotifyToken() (string, error) {
	client := resty.New()
	client.SetTimeout(10 * time.Second)

	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetBasicAuth(clientID, clientSecret).
		SetFormData(map[string]string{
			"grant_type": "client_credentials",
		}).
		Post("https://accounts.spotify.com/api/token")

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
	accessToken, err := getSpotifyToken()
	if err != nil {
		log.Fatalf("Failed to get access token: %v", err)
	}

	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+accessToken).
		Get("https://api.spotify.com/v1/me")

	if err != nil {
		log.Fatalf("Error making request to Spotify API: %v", err)
	}

	if resp.IsError() {
		log.Fatalf("Invalid token or request: %s", resp.String())
	}

	fmt.Println("Spotify Profile Response: ", resp)
}

func main() {
	testSpotifyToken()
	fmt.Println(clientID, clientSecret)
}
*/