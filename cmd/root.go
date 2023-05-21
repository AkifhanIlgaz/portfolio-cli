package cmd

import (
	"fmt"
	"os"

	"github.com/AkifhanIlgaz/portfolio/db"
	"github.com/AkifhanIlgaz/portfolio/price"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "portfolio",
	Short: "Show total balance and assets",
	Run: func(cmd *cobra.Command, args []string) {
		convert, _ := cmd.Flags().GetBool("try")
		if convert {
			tryPrice := price.TRY()
			balance := 0.

			for _, asset := range db.AllAssets() {
				assetPrice := price.Crypto(asset.Name).GetPrice()
				balance += assetPrice * asset.Balance
				fmt.Println(asset, assetPrice*tryPrice)
			}

			fmt.Println(balance * tryPrice)
			return
		} else {
			balance := 0.

			for _, asset := range db.AllAssets() {
				assetPrice := price.Crypto(asset.Name).GetPrice()
				balance += assetPrice * asset.Balance
				fmt.Println(asset, assetPrice)
			}

			fmt.Println(balance)
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("try", "t", false, "Convert to TRY")

}
