/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package pocketclient

type ItemList struct {
	Status int             `json:"status"`
	List   map[string]Item `json:"list"`
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

func (pocketClient PocketClient) Get(params *GetParams) (*ItemList, error) {
	data := struct {
		ConsumerKey string `json:"consumer_key"`
		AccessToken string `json:"access_token"`
		*GetParams
	}{
		ConsumerKey: pocketClient.ConsumerKey,
		AccessToken: pocketClient.accessToken,
		GetParams:   params,
	}

	itemList := ItemList{}
	err := postJSON("https://getpocket.com/v3/get", data, &itemList)
	if err != nil {
		return nil, err
	}

	return &itemList, nil
}
