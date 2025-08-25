/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/illidaris/aphroditecli/internal/encrypts"
	"github.com/spf13/cobra"
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "A base64 tool",
	Long: `A base64 tool. For example:

编码：
./aphroditecli base64 "{RAW1}" "{RAW2}"
解码：
./aphroditecli base64 -R "{EN1}" "{EN2}"
`,
	Run: func(cmd *cobra.Command, args []string) {
		if reverse {
			encrypts.Base64Decode(args...)
		} else {
			encrypts.Base64Encode(args...)
		}
	},
}

func init() {
	rootCmd.AddCommand(base64Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// base64Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// base64Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
