package client

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type PocketClient struct {
	consumerKey string
	accessToken string
	httpClient  http.Client
	redirectURI string
}

func NewPocketClient(consumerKey string) *PocketClient {
	return &PocketClient{
		consumerKey: consumerKey,
		httpClient:  http.Client{Timeout: 20 * time.Second}, // nolint
		redirectURI: "http://localhost:8000",
	}
}

func newPocketRequest(path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://getpocket.com/v3%s", path), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}
