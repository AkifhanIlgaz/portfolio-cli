package db

import (
	"encoding/binary"
	"fmt"
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

type TxType int

const (
	Buy TxType = iota
	Sell
)

// Key will be sequential ID
type Transaction struct {
	Name   string
	Date   time.Time
	Type   TxType
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
	var transactions []Transaction

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(txBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tx := Deserialize[Transaction](v)
			if tx.Name == asset {
				transactions = append(transactions, tx)
			}
		}
		return nil
	})

	return transactions

}

func CreateTransaction(newTx Transaction) (int, error) {
	var id int

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(txBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, Serialize(newTx))
	})

	if err != nil {
		return -1, err
	}

	return id, nil
}

func EditTransaction(id int, newTx Transaction) error {
	// newTx must be same asset as oldTx. If it's not, don't edit
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(txBucket)
		return b.Put(itob(id), Serialize(newTx))
	})

	if err != nil {
		return err
	}

	return nil
}

func DeleteTransaction(id int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(txBucket)
		_, err := GetTransaction(id)
		if err != nil {
			return fmt.Errorf("tx #%v doesn't exist", id)
		}

		return b.Delete(itob(id))
	})

	return err
}

func GetTransaction(id int) (Transaction, error) {
	var myTx Transaction

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(txBucket)

		if rawTx := b.Get(itob(id)); rawTx != nil {
			myTx = Deserialize[Transaction](rawTx)
			return nil
		} else {
			return fmt.Errorf("Cannot found transaction: %v", id)
		}
	})

	if err != nil {
		return Transaction{}, err
	}

	return myTx, nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
