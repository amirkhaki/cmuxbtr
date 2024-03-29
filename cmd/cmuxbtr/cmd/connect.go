package cmd

import (
	"context"
	"fmt"
	"github.com/amirkhaki/cmuxbtr/store"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func connect(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	fmt.Println("add called, args are:")
	id, err := cmd.Flags().GetInt("id")
	checkErr(err)
	url, err := cmd.Flags().GetString("url")
	checkErr(err)
	tax, err := cmd.Flags().GetFloat64("tax")
	checkErr(err)
	wage, err := cmd.Flags().GetFloat64("wage")
	checkErr(err)
	prdct := Product{Url: url, Tax: tax, Wage: wage}
	prdctByte, err := store.Encode(prdct)
	checkErr(err)
	err = store.Storage.Set(ctx, []byte(strconv.Itoa(id)), prdctByte)
	checkErr(err)
	fmt.Println("connected successfully!")
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
	connectCmd.Flags().Float64P("tax", "t", 0, "tax")
	connectCmd.Flags().Float64P("wage", "w", 0, "wage")
	connectCmd.MarkFlagRequired("url")
	connectCmd.Flags().Int("id", 0, "Product ID in your ecommerce")
	connectCmd.MarkFlagRequired("id")
}
