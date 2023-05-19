package cmd

import (
	"fmt"

	"github.com/AkifhanIlgaz/portfolio/db"
	"github.com/spf13/cobra"
)

var cryptoCmd = &cobra.Command{
	Use:   "crypto",
	Short: "Show crypto assets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(db.AllCryptoAssets())
	},
}

func init() {
	assetCmd.AddCommand(cryptoCmd)
}
