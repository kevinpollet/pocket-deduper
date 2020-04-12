package client

import (
	"bytes"
	"encoding/json"
	"time"
)

type ModifyAction struct {
	Action string `json:"action"`
	ItemID string `json:"item_id,omitempty"`
	Time   string `json:"time,omitempty"`
}

type ModifyResponse struct {
	Status        int           `json:"status"`
	ActionResults []interface{} `json:"action_results"`
	ActionErrors  []interface{} `json:"action_errors"`
}

type modifyRequest struct {
	ConsumerKey string         `json:"consumer_key"`
	AccessToken string         `json:"access_token"`
	Actions     []ModifyAction `json:"actions"`
}

func NewDeleteAction(itemID string) *ModifyAction {
	return &ModifyAction{
		Action: "delete",
		ItemID: itemID,
		Time:   string(time.Now().Unix()),
	}
}

func (p *PocketClient) Modify(actions []ModifyAction) (*ModifyResponse, error) {
	body := &modifyRequest{p.consumerKey, p.accessToken, actions}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := newPocketRequest("/send", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}

	res, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	modifyResponse := &ModifyResponse{}
	if err = json.NewDecoder(res.Body).Decode(modifyResponse); err != nil {
		return nil, err
	}

	return modifyResponse, nil
}
