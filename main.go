package main

import (
	"fmt"
	"kevinpollet/pocket-remove-duplicates/pkg/pocketclient"
	"log"
	"os"
)

func main() {
	pocketClient := pocketclient.PocketClient{
		ConsumerKey: os.Getenv("CONSUMER_KEY"),
	}

	err := pocketClient.Authorize()
	if err != nil {
		log.Fatal(err)
	}

	items, err := pocketClient.Get()
	if err != nil {
		log.Fatal(err)
	}

	d := make(map[string]*pocketclient.Item, 0)

	for _, item := range *items {
		existing := d[item.ResolvedURL]
		if existing == nil {
			d[item.ResolvedURL] = &item
		} else {
			fmt.Printf("--> Duplicate: %s\n", item.ResolvedTitle)
		}
	}
}
