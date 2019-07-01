/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package pocket

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
	Status int  `json:"status"`
	Item   Item `json:"item"`
}

func (client *PocketClient) Add(params *AddParams) (*AddResponse, error) {
	body := struct {
		ConsumerKey string `json:"consumer_key"`
		AccessToken string `json:"access_token"`
		*AddParams
	}{
		client.ConsumerKey,
		client.accessToken,
		params,
	}

	res, err := resty.
		R().
		SetResult(&AddResponse{}).
		SetBody(&body).
		Post(client.resolveURL("/add"))

	if err != nil {
		return nil, err
	}

	return res.Result().(*AddResponse), nil
}
