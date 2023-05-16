package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AkifhanIlgaz/portfolio-cli/cmd"
	"github.com/AkifhanIlgaz/portfolio-cli/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "portfolio.db")
	must(db.Init(dbPath))

	cmd.Execute()
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
