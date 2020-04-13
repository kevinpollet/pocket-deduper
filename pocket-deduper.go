package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kevinpollet/pocket-deduper/client"
	"github.com/kevinpollet/pocket-deduper/deduper"
)

const usage = `pocket-deduper [options]

Options:
-consumerKey  Pocket API consumer key.
-dryRun       Print duplicate items without removing them from Pocket.
-help         Prints this text.
`

var (
	consumerKey = flag.String("consumerKey", "", "")
	dryRun      = flag.Bool("dryRun", false, "")
)

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

	duplicateItems := deduper.GetDuplicateItems(res.List)

	if len(duplicateItems) == 0 {
		fmt.Println("\n✔ No duplicate items")
		return
	}

	deleteItemActions := make([]client.ModifyAction, 0)

	fmt.Println("\nDuplicate items:")

	for _, item := range duplicateItems {
		deleteItemActions = append(deleteItemActions, *client.NewDeleteAction(item.ItemID))
		fmt.Printf("● %s\n", item.ResolvedTitle)
	}

	if !*dryRun {
		_, err = pocketClient.Modify(deleteItemActions)
		if err != nil {
			log.Fatal(err)
		}
	}
}
