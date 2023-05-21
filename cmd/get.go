package cmd

import (
	"fmt"
	"strconv"

	"github.com/AkifhanIlgaz/portfolio/db"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Return the given transaction",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Invalid ID: #%v\n", arg)
				continue
			}
			tx, err := db.GetTransaction(id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			tbl := table.New("ID", "Asset", "Type", "Amount", "Price", "Date")
			tbl.AddRow(fmt.Sprintf("#%v", tx.ID), tx.Asset, tx.Type, tx.Amount, tx.Price, tx.Date.Format("2006-01-02 15:04"))
			tbl.Print()
		}
	},
}

func init() {
	txCmd.AddCommand(getCmd)
}
