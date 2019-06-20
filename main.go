package main

import (
	"fmt"
	"kevinpollet/pocket-remove-duplicates/internal/pkg/pocketclient"
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

	fmt.Println(pocketClient)
}
