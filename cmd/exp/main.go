package main

import (
	"path/filepath"

	"github.com/AkifhanIlgaz/portfolio-cli/db"
	"github.com/mitchellh/go-homedir"
)

func main() {

	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "portfolio.db")
	db.Init(dbPath)

	// oldTx := db.Transaction{
	// 	Name:   "Mumu",
	// 	Amount: 15,
	// 	Price:  1,
	// }

	// db.CreateTransaction(oldTx)

	// db.AllTransactions()

	db.GetTransaction(13)

	newTx := db.Transaction{
		Name:   "Bloktopia",
		Amount: 13,
		Price:  1,
	}

	db.EditTransaction(13, newTx)

	db.GetTransaction(13)

	db.AllTransactions()
}
