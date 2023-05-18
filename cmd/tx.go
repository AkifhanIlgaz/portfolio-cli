package cmd

import (
	"fmt"

	"github.com/AkifhanIlgaz/portfolio/db"
	"github.com/spf13/cobra"
)

var txCmd = &cobra.Command{
	Use:   "tx",
	Short: "Show all transactions",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Show in table
		// TODO: Validate asset name ?

		if len(args) == 0 {
			for _, tx := range db.AllTransactions() {
				fmt.Println(tx)
			}
			return
		}

		for _, asset := range args {
			for _, tx := range db.AllTransactionsOfAsset(asset) {
				fmt.Println(tx)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(txCmd)
}
