/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/AkifhanIlgaz/portfolio/price"
	"github.com/spf13/cobra"
)

// priceCmd represents the price command
var priceCmd = &cobra.Command{
	Use:   "price",
	Short: "Show the current price of currency",

	Run: func(cmd *cobra.Command, args []string) {
		convert, _ := cmd.Flags().GetBool("try")

		for _, currency := range args {

			if convert {
				tryPrice := price.TRY()
				currencyPrice := price.Crypto(currency).GetPrice()
				fmt.Println(currency, currencyPrice*tryPrice)
			} else {
				currencyPrice := price.Crypto(currency).GetPrice()
				fmt.Println(currency, currencyPrice)

			}

		}

	},
}

func init() {
	rootCmd.AddCommand(priceCmd)

	priceCmd.Flags().Bool("try", false, "convert to TRY")
}
