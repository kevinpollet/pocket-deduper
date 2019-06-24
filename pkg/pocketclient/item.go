/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package pocketclient

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
