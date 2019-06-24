/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package pocketclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type PocketClient struct {
	ConsumerKey string
	username    string
	accessToken string
}
