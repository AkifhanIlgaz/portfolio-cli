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

}
