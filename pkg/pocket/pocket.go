/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package pocket

const DefaultBaseURL = "https://getpocket.com/v3"

type PocketClient struct {
	BaseURL     string
	ConsumerKey string
	username    string
	accessToken string
}

func (client *PocketClient) resolveURL(path string) string {
	if client.BaseURL != "" {
		return client.BaseURL + path
	}
	return DefaultBaseURL + path
}
