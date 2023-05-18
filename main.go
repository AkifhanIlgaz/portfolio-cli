package main

import (
	"log"
	"path/filepath"

	"github.com/AkifhanIlgaz/portfolio/cmd"
	"github.com/AkifhanIlgaz/portfolio/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "portfolio.db")
	err := db.Init(dbPath)
	if err != nil {
		log.Fatal("Unable to connect BoltDB")
	}

	cmd.Execute()
}
