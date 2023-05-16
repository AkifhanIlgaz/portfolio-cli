package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteAssetCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an asset",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

func init() {
	assetCmd.AddCommand(deleteAssetCmd)
}
