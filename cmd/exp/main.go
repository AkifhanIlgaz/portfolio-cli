package main

import (
	"fmt"

	"github.com/AkifhanIlgaz/portfolio/price"
)

func main() {

	// home, _ := homedir.Dir()
	// dbPath := filepath.Join(home, "portfolio.db")
	// db.Init(dbPath)

	fmt.Println(price.Crypto("bitcoin", "ethereum", "    stafi", "Pegaxy Stone"))

	fmt.Println(price.TRY())
}
