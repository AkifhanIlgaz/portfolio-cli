package db

import (
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

func Init(dbPath string) error {
	var err error

	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(assetBucket)
		return err
	})
	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(txBucket)
		return err
	})
	if err != nil {
		return err
	}

	return nil
}

func Serialize[T Transaction | Asset](data T) []byte {
	serialized, _ := json.Marshal(data)
	return serialized
}

func Deserialize[T Transaction | Asset](data []byte) T {
	var t T

	json.Unmarshal(data, &t)

	return t
}
