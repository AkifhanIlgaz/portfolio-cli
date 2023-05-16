package db

import "time"

var txBucket = []byte("transactions")

/*
	Create  Interactive
	Delete
	Edit	Interactive
	Get
	All
	AllTransactionForSameType ?
*/

type txType int

const (
	buy = iota
	sell
)

// Key will be sequential ID
type Transaction struct {
	Name   string
	Date   time.Time `json:"date"`
	Type   txType    `json:"type"`
	Amount float64
	Price  float64
}

func AllTransactions() []Transaction {
	panic("Return all transactions")
}

func AllTransactionsOfAsset(asset string) []Transaction {
	panic("Return all transactions of specified asset")
}

// Interactive ( Ask user what is the asset,price etc.)
func CreateTransaction(tx Transaction) error {
	panic("Create new transaction")
}

// Interactive
func EditTransaction(id int) error {
	panic("Edit existing transaction. Error if tx doesn't exist")
}

func DeleteTransaction(id int) error {
	panic("Delete the transaction")
}

func GetTransaction(id int) (Transaction, error) {
	panic("Return transaction by id")
}
