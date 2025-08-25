/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/illidaris/aphroditecli/internal/encrypts"
	"github.com/spf13/cobra"
)

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "A sign tool",
	Long: `A sign tool. For example:

./aphroditecli sign --secret "{SECRET}" "{RAW1}" "{RAW2}"
`,
	Run: func(cmd *cobra.Command, args []string) {
		encrypts.Sign(secret, args...)
	},
}

func init() {
	rootCmd.AddCommand(signCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
