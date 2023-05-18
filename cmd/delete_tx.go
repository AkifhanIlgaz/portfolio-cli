package cmd

import (
	"fmt"
	"strconv"

	"github.com/AkifhanIlgaz/portfolio/db"
	"github.com/spf13/cobra"
)

var deleteTxCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a transaction",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("cannot delete tx #%v. Invalid id", arg)
				continue
			}
			err = db.DeleteTransaction(id)
			if err != nil {
				fmt.Println("Cannot delete:", err)
				return
			}
			fmt.Printf("Deleted tx #%v \n", id)
		}
	},
}

func init() {
	txCmd.AddCommand(deleteTxCmd)
}
