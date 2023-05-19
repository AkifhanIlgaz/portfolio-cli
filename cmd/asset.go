package cmd

import (
	"fmt"

	"github.com/AkifhanIlgaz/portfolio/db"
	"github.com/spf13/cobra"
)

var assetCmd = &cobra.Command{
	Use:   "asset",
	Short: "Show all assets",
	Run: func(cmd *cobra.Command, args []string) {
		for _, asset := range db.AllAssets() {
			fmt.Println(asset)
		}
	},
}

func init() {
	rootCmd.AddCommand(assetCmd)

}
