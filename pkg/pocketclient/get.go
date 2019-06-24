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

type Item struct {
	ID            string `json:"item_id"`
	ResolvedID    string `json:"resolved_id"`
	GivenURL      string `json:"given_url"`
	ResolvedURL   string `json:"resolved_url"`
	GivenTitle    string `json:"given_title"`
	ResolvedTitle string `json:"resolved_title"`
	Favorite      string `json:"favorite"`
	Status        string `json:"status"`
}

type GetResponse struct {
	Status int             `json:"status"`
	List   map[string]Item `json:"list"`
}

func (client *PocketClient) Get(params *GetParams) (*GetResponse, error) {
	data := struct {
		ConsumerKey string `json:"consumer_key"`
		AccessToken string `json:"access_token"`
		*GetParams
	}{
		ConsumerKey: client.ConsumerKey,
		AccessToken: client.accessToken,
		GetParams:   params,
	}

	res, err := resty.
    R().
    SetResult(&GetResponse{}).
		SetBody(&data).
		Post("https://getpocket.com/v3/get")

	if err != nil {
		return nil, err
	}

	return res.Result().(*GetResponse), nil
}
