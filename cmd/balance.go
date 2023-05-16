package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Show total balance",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("balance called")
	},
}

func init() {
	rootCmd.AddCommand(balanceCmd)
}
