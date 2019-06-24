/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package pocketclient

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
  return "https://getpocket.com/v3" + path
}
