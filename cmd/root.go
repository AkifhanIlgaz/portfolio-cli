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
		balance := 0.

		for _, asset := range db.AllAssets() {
			assetPrice := price.Crypto(asset.Name)
			fmt.Println(assetPrice)
			balance += assetPrice[asset.Name]["usd"] * asset.Balance
			fmt.Println(asset)
		}

		fmt.Println(balance)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
