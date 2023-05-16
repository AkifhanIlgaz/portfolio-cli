package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/AkifhanIlgaz/portfolio-cli/db"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new transaction",
	Run: func(cmd *cobra.Command, args []string) {
		tx := getTxInformationFromUser()
		db.CreateTransaction(tx)
	},
}

var scanner = bufio.NewScanner(os.Stdin)

func init() {
	txCmd.AddCommand(createCmd)
}

func getTxInformationFromUser() db.Transaction {
	return db.Transaction{
		Name:   getAssetName(),
		Date:   getDate(),
		Type:   getType(),
		Amount: getAmount(),
		Price:  getPrice(),
	}
}

func getAssetName() string {
	fmt.Print("> Asset Name: ")
	scanner.Scan()
	return strings.ToTitle(scanner.Text())
}

func getType() db.TxType {
	var txType db.TxType

	fmt.Print("> Transaction Type: ")
	scanner.Scan()
	switch scanner.Text() {
	case "sell":
		txType = db.Sell
	case "buy":
		txType = db.Buy
	}

	return txType
}

func getAmount() float64 {
	fmt.Print("> Amount: ")
	scanner.Scan()
	amount, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		log.Fatal("Unable to parse amount. Please enter valid amount")
	}

	return amount
}

func getPrice() float64 {
	fmt.Print("> Price: ")
	scanner.Scan()
	amount, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		log.Fatal("Unable to parse amount. Please enter valid price")
	}

	return amount
}

func getDate() time.Time {
	fmt.Print("> Date: ")
	scanner.Scan()
	if date := scanner.Text(); date == "" {
		return time.Now()
	}

	// Parse given time || DD-MM-YYYY

	return time.Now()
}
