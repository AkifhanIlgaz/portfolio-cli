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
		fmt.Println(db.AllTransactions())
	},
}

func init() {
	rootCmd.AddCommand(txCmd)
}
