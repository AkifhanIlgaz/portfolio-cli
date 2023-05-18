package db

import "github.com/boltdb/bolt"

var assetBucket = []byte("assets")

// Name will be key in DB
type Asset struct {
	Name        string
	Balance     float64
	AvgBuyPrice float64
	Type        string
}

func (a Asset) IsFiat() bool {
	switch a.Type {
	case "TRY", "BUSD", "USDT", "USDC", "DAI":
		return true
	}
	return false
}

func AllAssets() []Asset {
	panic("Return all assets that user has")
}

// Filter

func AllCryptoAssets() []Asset {
	var assets []Asset

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(assetBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			asset := Deserialize[Asset](v)
			if !asset.IsFiat() {
				assets = append(assets, asset)
			}
		}

		return nil
	})

	return assets
}

func AllFiatAssets() []Asset {
	var assets []Asset

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(assetBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			asset := Deserialize[Asset](v)
			if asset.IsFiat() {
				assets = append(assets, asset)
			}
		}

		return nil
	})

	return assets
}

func DeleteAsset(asset string) error {
	panic("Delete given asset. Return error if asset doesn't exist")
}

func UpdateAsset(asset string) error {
	panic(`
	Filter transactions by the name of given asset 
	Calculate balance and average buy price
	Delete if there is no transaction for this asset
	`)
}
