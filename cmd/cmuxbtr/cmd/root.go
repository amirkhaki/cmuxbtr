/*
Copyright © 2022 Amir Khaki

*/
package cmd

import (
	"github.com/amirkhaki/cmuxbtr/config"
	"github.com/amirkhaki/cmuxbtr/store"
	"os"
	"fmt"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cmuxbtr",
	Short: "cmuxbtr is a cli application which connects amaxon to your ecommerce",
	Long:  `Have fun :)`,
}

var cfg *config.Config
func init() {
	cfg, err := config.New()
	checkErr(err)
	err = store.Connect(cfg)
	fmt.Println("root.go, line 27")
	checkErr(err)
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
