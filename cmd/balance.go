package cmd

import (
	"fmt"

	"github.com/AkifhanIlgaz/portfolio/db"
	"github.com/AkifhanIlgaz/portfolio/price"
	"github.com/spf13/cobra"
)

var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Show total balance",
	Run: func(cmd *cobra.Command, args []string) {
		balance := 0.
		for _, asset := range db.AllAssets() {
			p := price.Crypto(asset.Name)
			balance += p[asset.Name].Usd * asset.Balance
		}
		fmt.Println(balance)
	},
}

func init() {
	rootCmd.AddCommand(balanceCmd)
}
