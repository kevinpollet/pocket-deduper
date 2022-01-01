package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

// GetParams represents the GET parameters sent to the Pocket API.
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

// Item represents a Pocket Item.
type Item struct {
	ItemID        string                 `json:"item_id"`
	ResolvedID    string                 `json:"resolved_id"`
	GivenURL      string                 `json:"given_url"`
	ResolvedURL   string                 `json:"resolved_url"`
	GivenTitle    string                 `json:"given_title"`
	ResolvedTitle string                 `json:"resolved_title"`
	Favorite      string                 `json:"favorite"`
	Status        string                 `json:"status"`
	Excerpt       string                 `json:"excerpt,omitempty"`
	IsArticle     string                 `json:"is_article"`
	HasImage      string                 `json:"has_image"`
	HasVideo      string                 `json:"has_video"`
	WordCount     string                 `json:"word_count"`
	Tags          map[string]interface{} `json:"tags"`
	Authors       map[string]interface{} `json:"authors"`
	Images        map[string]interface{} `json:"images"`
	Videos        map[string]interface{} `json:"videos"`
}

// GetResponse represents the GET response receives by the Pocket API.
type GetResponse struct {
	Status int             `json:"status"`
	List   map[string]Item `json:"list"`
}

type getRequest struct {
	GetParams

	ConsumerKey string `json:"consumer_key"`
	AccessToken string `json:"access_token"`
}

// Get calls the get endpoint of the Pocket API.
// See https://getpocket.com/developer/docs/v3/retrieve
func (p *PocketClient) Get(ctx context.Context, params GetParams) (GetResponse, error) {
	body := &getRequest{
		GetParams:   params,
		ConsumerKey: p.consumerKey,
		AccessToken: p.accessToken,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return GetResponse{}, fmt.Errorf("marshal: %w", err)
	}

	req, err := newRequest(ctx, "/get", bytes.NewReader(bodyBytes))
	if err != nil {
		return GetResponse{}, err
	}

	res, err := p.httpClient.Do(req)
	if err != nil {
		return GetResponse{}, fmt.Errorf("do request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode/100 != 2 {
		return GetResponse{}, fmt.Errorf("status code: %d", res.StatusCode)
	}

	var getResponse GetResponse
	if err = json.NewDecoder(res.Body).Decode(&getResponse); err != nil {
		return GetResponse{}, fmt.Errorf("decode response: %w", err)
	}

	return getResponse, nil
}
