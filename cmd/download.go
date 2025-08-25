/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/illidaris/aphroditecli/internal/download"
	"github.com/spf13/cobra"
)

var (
	needdir bool
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "A download tool",
	Long: `A download tool. For example:

./aphroditecli download --needdir --out ./ "{URL1}" "{URL2}" "{URL3}"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		download.Download(out, needdir, args...)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.PersistentFlags().BoolVar(&needdir, "needdir", false, "will create url path dir")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
