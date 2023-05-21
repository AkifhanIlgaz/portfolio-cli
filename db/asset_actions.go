package db

import (
	"fmt"

	"github.com/boltdb/bolt"
)

var assetBucket = []byte("assets")

// Name will be key in DB
type Asset struct {
	Name    string
	Balance float64
}

func (a Asset) isFiat() bool {
	switch a.Name {
	case "try", "busd", "usdt", "usdc", "dai":
		return true
	}
	return false
}

func AllAssets() []Asset {
	var assets []Asset

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(assetBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			assets = append(assets, Deserialize[Asset](v))
		}

		return nil
	})

	return assets

}

func AllCryptoAssets() []Asset {
	return filterCrypto(AllAssets())
}

// Reduce filter functions to one. Anonymous functions ?

func filterCrypto(assets []Asset) []Asset {
	var filtered []Asset

	for _, asset := range assets {
		if !asset.isFiat() {
			filtered = append(filtered, asset)
		}
	}

	return filtered
}

func filterFiat(assets []Asset) []Asset {
	var filtered []Asset

	for _, asset := range assets {
		if asset.isFiat() {
			filtered = append(filtered, asset)
		}
	}

	return filtered
}

func AllFiatAssets() []Asset {
	return filterFiat(AllAssets())
}

func DeleteAsset(asset string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(assetBucket)
		return b.Delete([]byte(asset))
	})
}

func GetAsset(name string) (Asset, error) {
	var asset Asset

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(assetBucket)
		if a := b.Get([]byte(name)); a == nil {
			return fmt.Errorf("asset cannot found: %v", name)
		} else {
			asset = Deserialize[Asset](a)
			return nil
		}
	})

	return asset, err
}

func UpdateAsset(transaction Transaction) error {
	asset, err := GetAsset(transaction.Asset)
	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		switch transaction.Type {
		case "buy", "airdrop":
			asset.Balance += transaction.Amount
		case "sell":
			if asset.Balance < transaction.Amount {
				return fmt.Errorf("not enough balance")
			}
			asset.Balance -= transaction.Amount
			if asset.Balance <= 0 {
				return DeleteAsset(asset.Name)
			}
		}

		return nil
	})

	return err
}
