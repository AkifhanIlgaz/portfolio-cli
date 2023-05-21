package cmd

import (
	"fmt"
	"os"

	"github.com/AkifhanIlgaz/portfolio/db"
	"github.com/AkifhanIlgaz/portfolio/price"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "portfolio",
	Short: "Show total balance and assets",
	Run: func(cmd *cobra.Command, args []string) {
		convert, _ := cmd.Flags().GetBool("try")
		tbl := table.New("Asset", "Balance", "Price", "Value")
		fmt.Println(convert)
		if convert {
			tryPrice := price.TRY()
			balance := 0.

			for _, asset := range db.AllAssets() {
				assetPrice := price.Crypto(asset.Name).GetPrice()
				balance += assetPrice * asset.Balance
				tbl.AddRow(asset.Name, asset.Balance, fmt.Sprintf("%v₺", assetPrice*tryPrice), assetPrice*asset.Balance*tryPrice)
			}

			fmt.Printf("Total Balance: %v₺\n", balance*tryPrice)
		} else {
			balance := 0.

			for _, asset := range db.AllAssets() {
				assetPrice := price.Crypto(asset.Name).GetPrice()
				balance += assetPrice * asset.Balance
				tbl.AddRow(asset.Name, asset.Balance, fmt.Sprintf("%v$", assetPrice), assetPrice*asset.Balance)
			}

			fmt.Printf("Total Balance: %v$\n", balance)
		}

		tbl.Print()

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("try", "t", false, "Convert to TRY")

}
