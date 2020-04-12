package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kevinpollet/pocket-deduper/pkg/client"
)

const usage = `pocket-deduper [options]

Options:
-consumerKey  Pocket API consumer key.
-help         Prints this text.
`

var consumerKey = flag.String("consumerKey", "", "")

func main() {
	flag.Usage = func() {
		fmt.Print(usage)
		os.Exit(2) // nolint
	}
	flag.Parse()

	pocketClient := client.NewPocketClient(*consumerKey)

	if err := pocketClient.Authorize(); err != nil {
		log.Fatal(err)
	}

	res, err := pocketClient.Get(&client.GetParams{})
	if err != nil {
		log.Fatal(err)
	}

	itemSet := make(map[string]*client.Item)
	deleteItemActions := make([]client.ModifyAction, 0)

	for _, item := range res.List {
		item := item

		if existingItem := itemSet[item.ResolvedURL]; existingItem == nil {
			itemSet[item.ResolvedURL] = &item
		} else {
			deleteItemActions = append(deleteItemActions, *client.NewDeleteAction(item.ItemID))
			fmt.Printf("\n● Duplicate item: %s", item.ResolvedTitle)
		}
	}

	if len(deleteItemActions) == 0 {
		fmt.Println("\n✔ No duplicate items found")
	} else {
		_, err = pocketClient.Modify(deleteItemActions)
		if err != nil {
			log.Fatal(err)
		}
	}
}
