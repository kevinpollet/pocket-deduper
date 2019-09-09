/**
 * Copyright © 2019 kevinpollet <pollet.kevin@gmail.com>`
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE.md file.
 */

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kevinpollet/pocket-remove-duplicates/internal/pocket"
)

func main() {
	pocketClient := pocket.Client{
		ConsumerKey: os.Getenv("CONSUMER_KEY"),
	}

	err := pocketClient.Authorize()
	if err != nil {
		log.Fatal(err)
	}

	res, err := pocketClient.Get(&pocket.GetParams{})
	if err != nil {
		log.Fatal(err)
	}

	d := make(map[string]*pocket.Item, 0)

	for _, item := range res.List {
		existing := d[item.ResolvedURL]
		if existing == nil {
			d[item.ResolvedURL] = &item
		} else {
			fmt.Printf("--> Duplicate: %s\n", item.ResolvedTitle)
		}
	}
}
