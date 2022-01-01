package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// PocketClient represents the Pocket API client.
type PocketClient struct {
	consumerKey string
	accessToken string
	httpClient  http.Client
	redirectURI *url.URL
}

// New returns a new PocketClient instance.
func New(consumerKey string) *PocketClient {
	return &PocketClient{
		consumerKey: consumerKey,
		httpClient:  http.Client{Timeout: 20 * time.Second},
		redirectURI: mustParseURL("http://localhost:8000"),
	}
}

func newRequest(ctx context.Context, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("https://getpocket.com/v3%s", path), body)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	req.Header.Set("X-Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func mustParseURL(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	return u
}
