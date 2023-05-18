package cmd

import (
	"fmt"
	"strconv"

	"github.com/AkifhanIlgaz/portfolio-cli/db"
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
			fmt.Println(tx)
		}
	},
}

func init() {
	txCmd.AddCommand(getCmd)
}
