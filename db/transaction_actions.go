package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

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
	Date   time.Time
	Type   txType
	Amount float64
	Price  float64
}

func AllTransactions() ([]Transaction, error) {
	transactions := []Transaction{}

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(txBucket)

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			transaction := Deserialize[Transaction](v)

			transactions = append(transactions, transaction)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return transactions, nil

}

func AllTransactionsOfAsset(asset string) []Transaction {
	panic("Return all transactions of specified asset")
}

// Interactive ( Ask user what is the asset,price etc.)
func CreateTransaction(tx Transaction) error {
	// Return transaction ID
	mockTx := Transaction{
		Name:   "Sui",
		Date:   time.Now(),
		Type:   sell,
		Amount: 1500,
		Price:  1.30,
	}

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(txBucket)
		id64, _ := b.NextSequence()
		id := int(id64)
		key := itob(id)
		return b.Put(key, Serialize(mockTx))
	})

	if err != nil {
		return err
	}

	return nil

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

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
