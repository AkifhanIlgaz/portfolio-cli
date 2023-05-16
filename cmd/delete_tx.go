package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteTxCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a transaction",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

func init() {
	txCmd.AddCommand(deleteTxCmd)
}
