package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kevinpollet/pocket-deduper/client"
	"github.com/kevinpollet/pocket-deduper/deduper"
)

const usage = `pocket-deduper [options]

Options:
-consumer-key  Sets the Pocket API consumer key.
-dry-run       Prints duplicate items without removing them from Pocket.
-help          Prints this text.
`

var (
	dryRun      = flag.Bool("dry-run", false, "")
	consumerKey = flag.String("consumer-key", "", "")
)

func main() {
	flag.Usage = func() {
		fmt.Print(usage)
		os.Exit(2)
	}
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	pocketClient := client.New(*consumerKey)

	if err := pocketClient.Authorize(ctx, os.Stdout); err != nil {
		exit(cancel, "Error while authorizing the Pocket client: %v", err)
	}

	res, err := pocketClient.Get(ctx, client.GetParams{})
	if err != nil {
		exit(cancel, "Error while retrieving the Pocket items: %v", err)
	}

	duplicateItems := deduper.GetDuplicateItems(res.List)
	if len(duplicateItems) == 0 {
		fmt.Println("\n✔ No duplicate items")
		return
	}

	fmt.Println("\nDuplicate items:")

	deleteItemActions := make([]client.Action, 0, len(duplicateItems))
	for _, item := range duplicateItems {
		deleteItemActions = append(deleteItemActions, *client.NewDeleteAction(item.ItemID))
		fmt.Printf("● %s\n", item.ResolvedTitle)
	}

	if !*dryRun {
		if _, err = pocketClient.Send(ctx, deleteItemActions); err != nil {
			exit(cancel, "Error while deleting duplicated items: %v", err)
		}
	}

	fmt.Println("\n✔ Duplicate items removed successfully")
}

func exit(cancelFunc context.CancelFunc, msg string, v interface{}) {
	cancelFunc()

	log.Printf(msg, v)
	os.Exit(1)
}
