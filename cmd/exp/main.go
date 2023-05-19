package main

import (
	"fmt"

	"github.com/AkifhanIlgaz/portfolio/price"
)

func main() {

	// home, _ := homedir.Dir()
	// dbPath := filepath.Join(home, "portfolio.db")
	// db.Init(dbPath)

	fmt.Println(price.Crypto("bitcoin", "ethereum", "stafi", "pegaxy-stone"))

	fmt.Println(price.TRY())
}
