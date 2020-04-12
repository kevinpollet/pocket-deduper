package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type requestTokenResponse struct {
	Code string `json:"code"`
}

type accessTokenResponse struct {
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

func (p *PocketClient) Authorize() error {
	u, err := url.Parse(p.redirectURI)
	if err != nil {
		return err
	}

	requestToken, err := p.getRequestToken()
	if err != nil {
		return err
	}

	syncChan := make(chan error)
	server := http.Server{
		Addr:         fmt.Sprintf("127.0.0.1:%s", u.Port()),
		ReadTimeout:  5 * time.Second,  // nolint
		WriteTimeout: 10 * time.Second, // nolint
	}

	server.Handler = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Add("Content-Type", "text/plain")
		rw.WriteHeader(http.StatusOK)
		fmt.Fprint(rw, "You can close the tab.")

		go server.Shutdown(context.Background()) // nolint
	})

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			syncChan <- err
		}

		close(syncChan)
	}()

	fmt.Printf("Open: https://getpocket.com/auth/authorize?request_token=%s&redirect_uri=%s\n", requestToken, p.redirectURI) // nolint

	err = <-syncChan
	if err != nil {
		return err
	}

	accessToken, err := p.getAccessToken(requestToken)
	if err != nil {
		return err
	}

	p.accessToken = accessToken

	return nil
}

func (p *PocketClient) getAccessToken(code string) (string, error) {
	body := fmt.Sprintf(`{"consumer_key":"%s","code":"%s"}`, p.consumerKey, code)

	req, err := newPocketRequest("/oauth/authorize", strings.NewReader(body))
	if err != nil {
		return "", err
	}

	res, err := p.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	accessTokenResponse := &accessTokenResponse{}
	if err = json.NewDecoder(res.Body).Decode(accessTokenResponse); err != nil {
		return "", err
	}

	return accessTokenResponse.AccessToken, nil
}

func (p *PocketClient) getRequestToken() (string, error) {
	body := fmt.Sprintf(`{"consumer_key":"%s","redirect_uri":"%s"}`, p.consumerKey, p.redirectURI)

	req, err := newPocketRequest("/oauth/request", strings.NewReader(body))
	if err != nil {
		return "", err
	}

	res, err := p.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	requestTokenResponse := &requestTokenResponse{}
	if err = json.NewDecoder(res.Body).Decode(requestTokenResponse); err != nil {
		return "", err
	}

	return requestTokenResponse.Code, nil
}
