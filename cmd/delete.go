/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/AkifhanIlgaz/portfolio/db"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete given assets",

	Run: func(cmd *cobra.Command, args []string) {
		for _, asset := range args {
			db.DeleteAsset(asset)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
