/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/illidaris/aphroditecli/internal/encrypts"
	"github.com/spf13/cobra"
)

// urlencodeCmd represents the urlencode command
var urlencodeCmd = &cobra.Command{
	Use:   "urlencode",
	Short: "A urlencode tool",
	Long: `A urlencode tool. For example:
编码：
./aphroditecli urlencode "{RAW1}" "{RAW2}"
解码：
./aphroditecli urlencode -R "{EN1}" "{EN2}"
`,
	Run: func(cmd *cobra.Command, args []string) {
		if reverse {
			encrypts.UrlDecode(args...)
		} else {
			encrypts.UrlEncode(args...)
		}
	},
}

func init() {
	rootCmd.AddCommand(urlencodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// urlencodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// urlencodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
