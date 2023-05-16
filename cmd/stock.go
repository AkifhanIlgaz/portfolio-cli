package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stockCmd = &cobra.Command{
	Use:   "stock",
	Short: "Show stocks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stock called")
	},
}

func init() {
	assetCmd.AddCommand(stockCmd)
}
