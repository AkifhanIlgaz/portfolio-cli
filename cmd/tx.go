package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var txCmd = &cobra.Command{
	Use:   "tx",
	Short: "Show all transactions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tx called")
	},
}

func init() {
	rootCmd.AddCommand(txCmd)
}
