/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package pocketclient

import (
	"context"
	"fmt"
  "net/http"
  "gopkg.in/resty.v1"
)

type requestTokenResponse struct {
  Code string `json:"code"`
}

type accessTokenResponse struct {
  Username string `json:"username"`
  AccessToken string `json:"access_token"`
}

func (client *PocketClient) Authorize() error {
	redirectURI := "http://localhost:8000"
	code, err := client.getRequestToken(redirectURI)
	if err != nil {
		return err
	}

	router := http.NewServeMux()
	server := http.Server{Addr: redirectURI[7:], Handler: router}
	router.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		res, _ := client.getAccessToken(code.Code)
		client.username = res.Username
		client.accessToken = res.AccessToken

		writer.WriteHeader(200)
		writer.Write([]byte("Close the tab"))
		go func() {
			server.Shutdown(context.Background())
		}()
	})

	fmt.Printf("Open: https://getpocket.com/auth/authorize?request_token=%s&redirect_uri=%s\n", code.Code, redirectURI)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (client *PocketClient) getAccessToken(code string) (*accessTokenResponse, error) {
  res, err := resty.
    R().
    SetHeader("X-Accept", "application/json").
    SetBody(&map[string]string {
      "consumer_key": client.ConsumerKey,
		  "code":         code,
    }).
    SetResult(&accessTokenResponse{}).
    Post(client.resolveURL("/oauth/authorize"))

	if err != nil {
		return nil, err
	}

	return res.Result().(*accessTokenResponse), nil
}

func (client *PocketClient) getRequestToken(redirectURI string) (*requestTokenResponse, error) {
  res, err := resty.
    R().
    SetHeader("X-Accept", "application/json").
    SetBody(&map[string]string {
      "consumer_key": client.ConsumerKey,
      "redirect_uri": redirectURI,
    }).
    SetResult(&requestTokenResponse{}).
    Post(client.resolveURL("/oauth/request"))

	if err != nil {
		return nil, err
	}
	return res.Result().(*requestTokenResponse), nil
}
