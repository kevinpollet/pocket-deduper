package deduper

import "github.com/kevinpollet/pocket-deduper/pkg/client"

func GetDuplicateItems(items map[string]client.Item) []client.Item {
	itemsByURL := make(map[string]bool)
	duplicates := make([]client.Item, 0)

	for _, item := range items {
		if itemsByURL[item.ResolvedURL] {
			duplicates = append(duplicates, item)
		}

		itemsByURL[item.ResolvedURL] = true
	}

	return duplicates
}
