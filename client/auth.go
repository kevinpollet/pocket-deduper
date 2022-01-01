package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

type requestTokenResponse struct {
	Code string `json:"code"`
}

type accessTokenResponse struct {
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

// Authorize negotiates an AccessToken with the Pocket API server.
func (p *PocketClient) Authorize(ctx context.Context, w io.Writer) error {
	requestToken, err := p.getRequestToken(ctx)
	if err != nil {
		return err
	}

	var once sync.Once

	errCh := make(chan error)
	closeCh := make(chan struct{})

	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "text/plain")
		rw.WriteHeader(http.StatusOK)

		_, _ = io.WriteString(rw, "Close the tab.")

		once.Do(func() { close(closeCh) })
	})

	srv := http.Server{
		Addr:         fmt.Sprintf("localhost:%s", p.redirectURI.Port()),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      handler,
	}

	go func() {
		_, _ = io.WriteString(w,
			fmt.Sprintf("Open: https://getpocket.com/auth/authorize?request_token=%s&redirect_uri=%s\n", requestToken, p.redirectURI))

		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("listen and serve: %w", err)
		}
	}()

	select {
	case <-closeCh:
	case err := <-errCh:
		return err

	case <-ctx.Done():
		if err := shutdownServer(context.Background(), &srv); err != nil {
			return err
		}

		return ctx.Err()
	}

	if err := shutdownServer(ctx, &srv); err != nil {
		return err
	}

	p.accessToken, err = p.getAccessToken(ctx, requestToken)
	if err != nil {
		return err
	}

	return nil
}

func (p *PocketClient) getRequestToken(ctx context.Context) (string, error) {
	body := fmt.Sprintf(`{"consumer_key":"%s","redirect_uri":"%s"}`, p.consumerKey, p.redirectURI)

	req, err := newRequest(ctx, "/oauth/request", strings.NewReader(body))
	if err != nil {
		return "", err
	}

	res, err := p.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("do request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode/100 != 2 {
		return "", fmt.Errorf("status code: %d", res.StatusCode)
	}

	var requestTokenRes requestTokenResponse
	if err = json.NewDecoder(res.Body).Decode(&requestTokenRes); err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}

	return requestTokenRes.Code, nil
}

func (p *PocketClient) getAccessToken(ctx context.Context, code string) (string, error) {
	body := fmt.Sprintf(`{"consumer_key":"%s","code":"%s"}`, p.consumerKey, code)

	req, err := newRequest(ctx, "/oauth/authorize", strings.NewReader(body))
	if err != nil {
		return "", err
	}

	res, err := p.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("do request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode/100 != 2 {
		return "", fmt.Errorf("status code: %d", res.StatusCode)
	}

	var accessTokenRes accessTokenResponse
	if err = json.NewDecoder(res.Body).Decode(&accessTokenRes); err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}

	return accessTokenRes.AccessToken, nil
}

func shutdownServer(ctx context.Context, srv *http.Server) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown server: %w", err)
	}

	return nil
}
