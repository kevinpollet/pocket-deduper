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

type GetParams struct {
	State       string `json:"state,omitempty"`
	Favorite    string `json:"favorite,omitempty"`
	Tag         string `json:"tag,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	Sort        string `json:"sort,omitempty"`
	DetailType  string `json:"detailType,omitempty"`
	Search      string `json:"search,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Since       int    `json:"since,omitempty"`
	Count       int    `json:"count,omitempty"`
	Offset      int    `json:"offset,omitempty"`
}

type GetResponse struct {
	Status int             `json:"status"`
	List   map[string]Item `json:"list"`
}

func (client *Client) Get(params *GetParams) (*GetResponse, error) {
	body := struct {
		ConsumerKey string `json:"consumer_key"`
		AccessToken string `json:"access_token"`
		*GetParams
	}{
		client.ConsumerKey,
		client.accessToken,
		params,
	}

	res, err := resty.
		R().
		SetResult(&GetResponse{}).
		SetBody(&body).
		Post("https://getpocket.com/v3/get")

	if err != nil {
		return nil, err
	}

	return res.Result().(*GetResponse), nil
}
