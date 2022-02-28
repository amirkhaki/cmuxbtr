/*
Copyright © 2022 Amir Khaki

*/
package cmd

import (
	"context"
	"fmt"
	"github.com/amirkhaki/cmuxbtr/store"

	"github.com/spf13/cobra"
)

func list(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	for key := range store.Storage.Keys(ctx) {
		val, err := store.Storage.Get(ctx, key)
		checkErr(err)
		var prdct Product
		err = store.Decode(val, &prdct)
		checkErr(err)
		fmt.Println(string(key), " ==> ", prdct)
	}
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list of connections",
	Long:  `list`,
	Run:   list,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
