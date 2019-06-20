package pocketclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type PocketClient struct {
	ConsumerKey string
	redirectURI string
}

func (pocketClient PocketClient) RedirectURI() string {
	if pocketClient.redirectURI == "" {
		return "http://localhost:8000"
	} else {
		return pocketClient.redirectURI
	}
}

func (pocketClient PocketClient) GetAuthorizeURL(code string) string {
	return fmt.Sprintf("https://getpocket.com/auth/authorize?request_token=%s&redirect_uri=%s\n", code, pocketClient.RedirectURI())
}

func (pocketClient PocketClient) GetAccessToken(code string) (string, string) {
	body, _ := json.Marshal(map[string]string{
		"consumer_key": pocketClient.ConsumerKey,
		"code":         code,
	})

	req, _ := http.NewRequest("POST", "https://getpocket.com/v3/oauth/authorize", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json; charset=UTF8")
	req.Header.Add("X-Accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	var responseBody map[string]interface{}
	json.NewDecoder(res.Body).Decode(&responseBody)

	return responseBody["username"].(string), responseBody["access_token"].(string)
}

func (pocketClient PocketClient) GetRequestToken() (string, error) {
	body, _ := json.Marshal(map[string]string{
		"consumer_key": pocketClient.ConsumerKey,
		"redirect_uri": pocketClient.RedirectURI(),
	})

	req, _ := http.NewRequest("POST", "https://getpocket.com/v3/oauth/request", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json; charset=UTF8")
	req.Header.Add("X-Accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	var responseBody map[string]interface{}
	json.NewDecoder(res.Body).Decode(&responseBody)

	return responseBody["code"].(string), nil
}
