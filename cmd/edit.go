package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/AkifhanIlgaz/portfolio/db"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit existing transaction",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please enter a transaction ID")
			return
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Please enter a valid transaction ID")
		}

		oldTx, err := db.GetTransaction(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		// print old transaction as row. So user can see old values
		fmt.Println(oldTx)

		oldTx.Asset = changeAsset(oldTx.Asset)
		oldTx.Date = changeDate(oldTx.Date)
		oldTx.Type = changeType(oldTx.Type)
		oldTx.Amount = changeAmount(oldTx.Amount)
		oldTx.Price = changePrice(oldTx.Price)

		db.EditTransaction(id, oldTx)
		db.UpdateAsset(oldTx)
		fmt.Printf("Successfully updated transaction #%v\n", id)
	},
}

func changeAsset(oldAsset string) string {
	fmt.Print("> Asset Name: ")
	scanner.Scan()

	asset := scanner.Text()
	if asset == "" {
		return oldAsset
	}

	return strings.TrimSpace(strings.ToLower(asset))
}

func changeType(oldType string) string {

	fmt.Print("> Transaction Type: ")
	scanner.Scan()
	if t := strings.TrimSpace(scanner.Text()); t != "" {
		return t
	}

	return oldType
}

func changeAmount(oldAmount float64) float64 {
	fmt.Print("> Amount: ")
	scanner.Scan()

	if amount := strings.TrimSpace(scanner.Text()); amount == "" {
		return oldAmount
	} else {
		amount, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			log.Fatal("Unable to parse amount. Please enter valid amount")
		}
		return amount
	}

}

func changePrice(oldPrice float64) float64 {
	fmt.Print("> Price: ")
	scanner.Scan()

	if price := strings.TrimSpace(scanner.Text()); price == "" {
		return oldPrice
	} else {
		price, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Fatal("Unable to parse price. Please enter valid price")
		}
		return price
	}

}

func changeDate(oldDate time.Time) time.Time {
	fmt.Print("> Date: ")
	scanner.Scan()
	if date := strings.TrimSpace(scanner.Text()); date == "" {
		return oldDate
	} else {
		newDate, err := time.Parse("01-02-2006", date)
		if err != nil {
			fmt.Println(err)
			return time.Time{}
		}
		return newDate
	}

}

func init() {
	txCmd.AddCommand(editCmd)
}
