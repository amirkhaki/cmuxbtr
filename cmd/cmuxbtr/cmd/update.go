/*
Copyright © 2022 Amir Khaki

*/
package cmd

import (
	"github.com/amirkhaki/cmuxbtr/scraper"
	"github.com/amirkhaki/cmuxbtr/poster"
	"github.com/amirkhaki/cmuxbtr/store"
	"fmt"
	"log"
	"strconv"
	"context"
	"github.com/spf13/cobra"
)

func updateOne(ctx context.Context, key []byte) error {
	val, err := store.Storage.Get(ctx, key)
	if err != nil {
		return fmt.Errorf("Could not get item by key %s: %w", key, err)
	}
	out, err := scraper.Scrape(ctx, string(val))
	checkErr(err)
	keyInt, err := strconv.Atoi(string(key))
	if err != nil {
		return fmt.Errorf("Could not convert key to int: %w", err)
	}
	err = poster.Post(ctx, cfg, keyInt, out)
	if err != nil {
		return fmt.Errorf("Could not post item with key %s: %w", key, err)
	}
	fmt.Println(out)
	return nil
}
// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `hi`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		checkErr(err)
		ctx := context.Background()
		if id != 0 {
			err = updateOne(ctx, []byte(strconv.Itoa(id)))
			checkErr(err)
			log.Printf("%d updated successfully")
			return
		}
		for key := range(store.Storage.Keys(ctx)) {
			log.Printf("updating %s", key)
			err = updateOne(ctx, key)
			checkErr(err)
		}
		log.Println("all items updated successfully!")

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	updateCmd.Flags().Int("id", 0, "Product id in your ecommerce")
}
