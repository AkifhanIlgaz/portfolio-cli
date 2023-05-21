package cmd

import (
	"fmt"

	"github.com/AkifhanIlgaz/portfolio/db"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var txCmd = &cobra.Command{
	Use:   "tx",
	Short: "Show all transactions",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Show in table

		tbl := table.New("ID", "Asset", "Type", "Amount", "Price", "Date")

		if len(args) == 0 {
			for _, tx := range db.AllTransactions() {
				tbl.AddRow(fmt.Sprintf("#%v", tx.ID), tx.Asset, tx.Type, tx.Amount, tx.Price, tx.Date.Format("2006-01-02 15:04"))
			}
			tbl.Print()
			return
		}

		for _, asset := range args {
			for _, tx := range db.AllTransactionsOfAsset(asset) {
				tbl.AddRow(fmt.Sprintf("#%v", tx.ID), tx.Asset, tx.Type, tx.Amount, tx.Price, tx.Date.Format("2006-01-02 15:04"))
			}
			tbl.Print()
		}

	},
}

func init() {
	rootCmd.AddCommand(txCmd)
}
