package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/AkifhanIlgaz/portfolio/db"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new transaction",
	Run: func(cmd *cobra.Command, args []string) {
		tx := getTxInformationFromUser()
		db.CreateTransaction(tx)
		db.UpdateAsset(tx)
	},
}

var scanner = bufio.NewScanner(os.Stdin)

func init() {
	txCmd.AddCommand(createCmd)
}

func getTxInformationFromUser() db.Transaction {
	return db.Transaction{
		Asset:  getAsset(),
		Date:   getDate(),
		Type:   getType(),
		Amount: getAmount(),
		Price:  getPrice(),
	}
}

func getAsset() string {
	fmt.Print("> Asset: ")
	scanner.Scan()

	asset := scanner.Text()
	if asset == "" {
		fmt.Println("Please enter an asset")
		return getAsset()
	}

	asset = strings.TrimSpace(asset)
	asset = strings.Join(strings.Fields(asset), "-")
	asset = strings.ToLower(asset)

	return asset
}

func getType() string {
	var txType string

	fmt.Print("> Transaction Type: ")
	scanner.Scan()
	t := strings.TrimSpace(strings.ToLower(scanner.Text()))
	switch t {
	case "sell", "buy", "airdrop":
		txType = t
	default:
		fmt.Println("Please enter valid transaction type")
		return getType()
	}

	return txType
}

func getAmount() float64 {
	fmt.Print("> Amount: ")
	scanner.Scan()
	amount, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
	if err != nil {
		fmt.Println("Please enter valid amount")
		return getAmount()
	}

	return amount
}

func getPrice() float64 {
	fmt.Print("> Price: ")
	scanner.Scan()
	amount, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
	if err != nil {
		fmt.Println("Please enter valid price")
		getPrice()
	}

	return amount
}

func getDate() time.Time {
	fmt.Print("> Date: ")
	scanner.Scan()
	if date := strings.TrimSpace(scanner.Text()); date == "" {
		return time.Now()
	} else {
		date, err := time.Parse("02-01-2006 15:04", date)
		if err != nil {
			fmt.Println(err)
			return time.Time{}
		}
		return date
	}

}
