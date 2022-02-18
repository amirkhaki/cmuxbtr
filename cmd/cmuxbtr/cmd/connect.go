package cmd

import (
	"fmt"

	"log"
	"github.com/spf13/cobra"
)
func checkErr( err error ) {
	if err != nil {
		log.Fatal(err)
	}
}
func connect(cmd *cobra.Command, args []string) {
	fmt.Println("add called, args are:")
	id, err := cmd.Flags().GetInt("id")
	checkErr(err)
	url, err := cmd.Flags().GetString("url")
	checkErr(err)
	fmt.Println(id, url)
}

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect a product in Amazon to a product in your ecommerce",
	Long: `Connect a product in Amazon to a product in your ecommerce,
	for example:
	cmuxbtr connect -u https://www.amazon.ae/dp/B09B1H2Q4R -i 12345`,
	Run: connect,
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.Flags().StringP("url", "u", "", "Product URL in Amazon")
	connectCmd.MarkFlagRequired("url")
	connectCmd.Flags().Int("id", 0, "Product ID in your ecommerce")
	connectCmd.MarkFlagRequired("id")
}
