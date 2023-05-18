package cmd

import (
	"fmt"

	"github.com/AkifhanIlgaz/portfolio-cli/db"
	"github.com/spf13/cobra"
)

var txCmd = &cobra.Command{
	Use:   "tx",
	Short: "Show all transactions",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Show in table
		// TODO: Validate asset name ?

		if len(args) == 0 {
			if txs, err := db.AllTransactions(); err == nil {
				for _, tx := range txs {
					fmt.Println(tx)
				}
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
