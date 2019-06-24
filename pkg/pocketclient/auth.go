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
)

func (pocketClient *PocketClient) Authorize() error {
	redirectURI := "http://localhost:8000"
	code, err := pocketClient.getRequestToken(redirectURI)
	if err != nil {
		return err
	}

	router := http.NewServeMux()
	server := http.Server{Addr: redirectURI[7:], Handler: router}
	router.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		username, accessToken, _ := pocketClient.getAccessToken(code)
		pocketClient.username = username
		pocketClient.accessToken = accessToken

		writer.WriteHeader(200)
		writer.Write([]byte("Close the tab"))
		go func() {
			server.Shutdown(context.Background())
		}()
	})

	fmt.Printf("Open: https://getpocket.com/auth/authorize?request_token=%s&redirect_uri=%s\n", code, redirectURI)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (pocketClient PocketClient) getAccessToken(code string) (string, string, error) {
	body := struct {
		Username    string `json:"username"`
		AccessToken string `json:"access_token"`
	}{}

	err := postJSON("https://getpocket.com/v3/oauth/authorize", map[string]string{
		"consumer_key": pocketClient.ConsumerKey,
		"code":         code,
	}, &body)
	if err != nil {
		return "", "", err
	}

	return body.Username, body.AccessToken, nil
}

func (pocketClient PocketClient) getRequestToken(redirectURI string) (string, error) {
	body := struct {
		Code string `json:"code"`
	}{}

	err := postJSON("https://getpocket.com/v3/oauth/request",
		map[string]string{
			"consumer_key": pocketClient.ConsumerKey,
			"redirect_uri": redirectURI,
		}, &body)

	if err != nil {
		return "", err
	}

	return body.Code, nil
}
