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

	"github.com/kevinpollet/pocket-list-dedupe/pocket"
	"github.com/spf13/cobra"
)

var (
	consumerKey string
	rootCmd     = &cobra.Command{
		Use:   "pocket-list-dedupe",
		Short: "Remove duplicate items in your Pocket reading list",
		Run: func(cmd *cobra.Command, args []string) {
			pocketClient := pocket.Client{ConsumerKey: consumerKey}
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
					fmt.Printf("--> Duplicate: %s/%s\n", item.ResolvedTitle, item.ResolvedURL)
				}
			}
		},
	}
)

func init() {
	rootCmd.Flags().StringVarP(&consumerKey, "consumerKey", "c", "", "Pocket application's Consumer Key")
	rootCmd.MarkFlagRequired("consumerKey")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
