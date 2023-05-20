/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// assetCmd represents the asset command
var assetCmd = &cobra.Command{
	Use: "asset",
}

func init() {
	rootCmd.AddCommand(assetCmd)
}
