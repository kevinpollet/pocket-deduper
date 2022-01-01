package deduper

import "github.com/kevinpollet/pocket-deduper/client"

// GetDuplicateItems returns the duplicated items in the given client.Item map.
func GetDuplicateItems(items map[string]client.Item) []client.Item {
	itemsByURL := make(map[string]struct{})

	var ret []client.Item

	for _, item := range items {
		if _, exists := itemsByURL[item.ResolvedURL]; exists {
			ret = append(ret, item)
		}

		itemsByURL[item.ResolvedURL] = struct{}{}
	}

	return ret
}
