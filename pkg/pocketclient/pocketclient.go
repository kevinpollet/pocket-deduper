/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package pocketclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type PocketClient struct {
	ConsumerKey string
	username    string
	accessToken string
}

func postJSON(url string, jsonBody interface{}, decodedBody interface{}) error {
	body, err := json.Marshal(jsonBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json; charset=UTF8")
	req.Header.Add("X-Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}

	return json.NewDecoder(res.Body).Decode(decodedBody)
}
