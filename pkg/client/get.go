/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package client

import (
	"bytes"
	"encoding/json"
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

type getRequest struct {
	ConsumerKey string `json:"consumer_key"`
	AccessToken string `json:"access_token"`
	*GetParams
}

func (p *PocketClient) Get(params *GetParams) (*GetResponse, error) {
	body := &getRequest{p.consumerKey, p.accessToken, params}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := newPocketRequest("/get", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}

	res, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	getResponse := &GetResponse{}
	if err = json.NewDecoder(res.Body).Decode(getResponse); err != nil {
		return nil, err
	}

	return getResponse, nil
}
