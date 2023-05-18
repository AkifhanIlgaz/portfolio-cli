package db

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var txBucket = []byte("transactions")

type TxType int

const (
	Buy TxType = iota
	Sell
	InvalidTxType
)

// Key will be sequential ID
type Transaction struct {
	Asset  string
	Date   time.Time
	Type   TxType
	Amount float64
	Price  float64
}

func AllTransactions() []Transaction {
	transactions := []Transaction{}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(txBucket)

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			transaction := Deserialize[Transaction](v)
			transactions = append(transactions, transaction)
		}

		return nil
	})

	return transactions

}

func AllTransactionsOfAsset(asset string) []Transaction {
	var transactions []Transaction

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(txBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tx := Deserialize[Transaction](v)
			if tx.Asset == asset {
				transactions = append(transactions, tx)
			}
		}
		return nil
	})

	return transactions

}

func GetAllTransactionsOfType(txType TxType) []Transaction {
	var transactions []Transaction

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(txBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tx := Deserialize[Transaction](v)
			if tx.Type == txType {
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
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(txBucket)
		return b.Put(itob(id), Serialize(newTx))
	})
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
		}

		return fmt.Errorf("cannot found transaction: %v", id)
	})

	return myTx, err
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
