package db

var assetBucket = []byte("assets")

/*
	AllAssets
	Crypto
	Stock
	Delete
*/

type AssetType int

const (
	Crypto = iota
	Stock
	Fiat
)

// Name will be key in DB
type Asset struct {
	Balance  float64
	AvgPrice float64
	Type     string
}

func AllAssets() []Asset {
	panic("Return all assets that user has")
}

func Crpyto() []Asset {
	panic("Return all cryptocurrencies that user has")
}

func Stocks() []Asset {
	panic("Return all stocks that user has")
}

func DeleteAsset(name string) error {
	panic("Delete given asset. Return error if asset doesn't exist")
}

func UpdateAsset(name string) error {
	panic(`
	Filter transactions by the name of given asset 
	Calculate balance and average buy price
	Delete if there is no transaction for this asset
	`)
}
