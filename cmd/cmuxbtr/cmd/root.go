/*
Copyright © 2022 Amir Khaki

*/
package cmd

import (
	"fmt"
	"github.com/amirkhaki/cmuxbtr/config"
	"github.com/amirkhaki/cmuxbtr/store"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cmuxbtr",
	Short: "cmuxbtr is a cli application which connects amaxon to your ecommerce",
	Long:  `Have fun :)`,
	Run: updateCmdFunc,
}

func init() {
	err := config.Parse()
	checkErr(err)
	err = store.Connect()
	fmt.Println("root.go, line 27")
	checkErr(err)
	rootCmd.Flags().Int("id", 0, "Product id in your ecommerce")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	defer store.Close()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
