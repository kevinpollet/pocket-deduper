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

	"github.com/kevinpollet/pocket-list-dedupe/pocket"
	"github.com/spf13/cobra"
)

var (
	consumerKey string
	rootCmd     = &cobra.Command{
		Use:   "pocket-list-dedupe",
		Short: "Remove duplicate items in your Pocket list",
		Run: func(cmd *cobra.Command, args []string) {
			pocketClient := pocket.Client{
				ConsumerKey: consumerKey,
			}

			if err := pocketClient.Authorize(); err != nil {
				log.Fatal(err)
			}

			res, err := pocketClient.Get(&pocket.GetParams{})
			if err != nil {
				log.Fatal(err)
			}

			itemSet := make(map[string]*pocket.Item, 0)
			deleteItemActions := make([]pocket.ModifyAction, 0)

			for _, item := range res.List {
				if existingItem := itemSet[item.ResolvedURL]; existingItem == nil {
					itemSet[item.ResolvedURL] = &item
				} else {
					deleteItemActions = append(deleteItemActions, *pocket.NewDeleteAction(item.ItemID))
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
		},
	}
)

func main() {
	rootCmd.Flags().StringVarP(&consumerKey, "consumerKey", "c", "", "Pocket application's Consumer Key")
	rootCmd.MarkFlagRequired("consumerKey")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
