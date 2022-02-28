/*
Copyright © 2022 Amir Khaki

*/
package cmd

import (
	"github.com/amirkhaki/cmuxbtr/store"

	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete connection",
	Long:  `delete`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
		ctx := context.Background()
		id, err := cmd.Flags().GetInt("id")
		checkErr(err)
		err = store.Storage.Delete(ctx, []byte(strconv.Itoa(id)))
		checkErr(err)
		fmt.Printf("%d deleted successfully", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().Int("id", 0, "Product ID in your ecommerce")
	deleteCmd.MarkFlagRequired("id")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
