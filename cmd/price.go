/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/AkifhanIlgaz/portfolio/price"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

// priceCmd represents the price command
var priceCmd = &cobra.Command{
	Use:   "price",
	Short: "Show the current price of currency",

	Run: func(cmd *cobra.Command, args []string) {
		convert, _ := cmd.Flags().GetBool("try")
		tbl := table.New("Asset", "Price")

		for _, currency := range args {

			if convert {
				tryPrice := price.TRY()
				currencyPrice := price.Crypto(currency).GetPrice()
				tbl.AddRow(currency, fmt.Sprintf("%v₺", currencyPrice*tryPrice))
			} else {
				currencyPrice := price.Crypto(currency).GetPrice()
				tbl.AddRow(currency, fmt.Sprintf("%v$", currencyPrice))
			}

		}

		tbl.Print()
	},
}

func init() {
	rootCmd.AddCommand(priceCmd)

	priceCmd.Flags().BoolP("try", "t", false, "convert to TRY")
}
