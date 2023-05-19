package db

import "github.com/boltdb/bolt"

var assetBucket = []byte("assets")

// Name will be key in DB
type Asset struct {
	Name        string
	Balance     float64
	AvgBuyPrice float64
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

func UpdateAsset(asset string) error {

	transactions := AllTransactionsOfAsset(asset)

	if len(transactions) == 0 {
		return DeleteAsset(asset)
	}

	var a = Asset{
		Name: asset,
	}

	buyAmount, totalMoneyToBuy := 0., 0.

	for _, transaction := range transactions {
		if transaction.Type == "buy" {
			buyAmount += transaction.Amount
			totalMoneyToBuy += transaction.Price * transaction.Amount
			a.Balance += transaction.Amount
		} else if transaction.Type == "airdrop" {
			a.Balance += transaction.Amount
		} else {
			a.Balance -= transaction.Amount
		}
	}

	a.AvgBuyPrice = totalMoneyToBuy / buyAmount

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(assetBucket)
		return b.Put([]byte(asset), Serialize(a))
	})

}
