/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package pocket

import (
	"time"

	"gopkg.in/resty.v1"
)

type ModifyResponse struct {
	Status        int           `json:"status"`
	ActionResults []interface{} `json:"action_results"`
	ActionErrors  []interface{} `json:"action_errors"`
}

type ModifyAction struct {
	Action string `json:"action"`
	ItemID string `json:"item_id,omitempty"`
	Time   string `json:"time,omitempty"`
}

func NewDeleteAction(itemID string) *ModifyAction {
	return &ModifyAction{
		Action: "delete",
		ItemID: itemID,
		Time:   string(time.Now().Unix()),
	}
}

func (client *Client) Modify(actions []ModifyAction) (*ModifyResponse, error) {
	body := struct {
		ConsumerKey string         `json:"consumer_key"`
		AccessToken string         `json:"access_token"`
		Actions     []ModifyAction `json:"actions"`
	}{
		client.ConsumerKey,
		client.accessToken,
		actions,
	}

	res, err := resty.
		R().
		SetHeader("X-Accept", "application/json").
		SetResult(ModifyResponse{}).
		SetBody(body).
		Post("https://getpocket.com/v3/send")

	if err != nil {
		return nil, err
	}
	return res.Result().(*ModifyResponse), nil
}
