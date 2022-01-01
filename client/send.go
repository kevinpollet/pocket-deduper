package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Action represents an Action sent to the Pocket API.
type Action struct {
	Type   string `json:"action"`
	ItemID string `json:"item_id,omitempty"`
	Time   string `json:"time,omitempty"`
}

// NewDeleteAction creates a delete action for the given itemID.
func NewDeleteAction(itemID string) *Action {
	return &Action{
		Type:   "delete",
		ItemID: itemID,
		Time:   strconv.Itoa(int(time.Now().Unix())),
	}
}

// SendResponse represents the send endpoint response.
type SendResponse struct {
	Status        int           `json:"status"`
	ActionResults []interface{} `json:"action_results"`
	ActionErrors  []interface{} `json:"action_errors"`
}

type sendRequest struct {
	ConsumerKey string   `json:"consumer_key"`
	AccessToken string   `json:"access_token"`
	Actions     []Action `json:"actions"`
}

// Send calls the send endpoint of the Pocket API.
// See https://getpocket.com/developer/docs/v3/modify
func (p *PocketClient) Send(ctx context.Context, actions []Action) (SendResponse, error) {
	body := &sendRequest{
		ConsumerKey: p.consumerKey,
		AccessToken: p.accessToken,
		Actions:     actions,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return SendResponse{}, fmt.Errorf("marshal: %w", err)
	}

	req, err := newRequest(ctx, "/send", bytes.NewReader(bodyBytes))
	if err != nil {
		return SendResponse{}, err
	}

	res, err := p.httpClient.Do(req)
	if err != nil {
		return SendResponse{}, fmt.Errorf("do request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode/100 != 2 {
		return SendResponse{}, fmt.Errorf("status code: %d", res.StatusCode)
	}

	var modifyResponse SendResponse
	if err = json.NewDecoder(res.Body).Decode(&modifyResponse); err != nil {
		return SendResponse{}, fmt.Errorf("decode response: %w", err)
	}

	return modifyResponse, nil
}
