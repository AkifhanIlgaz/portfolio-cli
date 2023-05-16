package main

import (
	"fmt"

	"github.com/AkifhanIlgaz/portfolio-cli/db"
)

func main() {

	asset := db.Asset{
		Type:     db.Crypto,
		Balance:  100,
		AvgPrice: 1,
	}

	x := db.Serialize(asset)
	fmt.Println(x)

	fmt.Println(db.Deserialize[db.Asset](x))
}
