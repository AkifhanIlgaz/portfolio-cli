package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var assetCmd = &cobra.Command{
	Use:   "asset",
	Short: "Show all assets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("asset called")
	},
}

func init() {
	rootCmd.AddCommand(assetCmd)

}
