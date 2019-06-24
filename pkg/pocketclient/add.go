/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package pocketclient

import (
  "gopkg.in/resty.v1"
)

type AddParams struct {
	URL     string `json:"url"`
	Title   string `json:"title,omitempty"`
	Tags    string `json:"tags,omitempty"`
	TweetID string `json:"tweet_id,omitempty"`
}

type AddResponse struct {
	Status int         `json:"status"`
	Item   interface{} `json:"item"`
}

func (client *PocketClient) Add(params *AddParams) (*AddResponse, error) {
	data := struct {
		ConsumerKey string `json:"consumer_key"`
		AccessToken string `json:"access_token"`
		*AddParams
	}{
		ConsumerKey: client.ConsumerKey,
		AccessToken: client.accessToken,
		AddParams:   params,
  }

  res, err := resty.R().
    SetResult(&AddResponse{}).
    SetBody(&data).
    Post("https://getpocket.com/v3/add")

	if err != nil {
		return nil, err
	}

	return res.Result().(*AddResponse), nil
}
