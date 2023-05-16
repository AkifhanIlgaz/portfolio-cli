package cmd

import (
	"github.com/AkifhanIlgaz/portfolio-cli/db"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new transaction",
	Run: func(cmd *cobra.Command, args []string) {
		db.CreateTransaction(db.Transaction{})
	},
}

func init() {
	txCmd.AddCommand(createCmd)
}
