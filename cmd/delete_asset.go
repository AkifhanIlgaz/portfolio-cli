package cmd

import (
	"github.com/AkifhanIlgaz/portfolio/db"
	"github.com/spf13/cobra"
)

var deleteAssetCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an asset",
	Run: func(cmd *cobra.Command, args []string) {
		for _, asset := range args {
			db.DeleteAsset(asset)
		}
	},
}

func init() {
	assetCmd.AddCommand(deleteAssetCmd)
}
