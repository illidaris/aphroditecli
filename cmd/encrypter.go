/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/illidaris/aphroditecli/internal/encrypts"
	"github.com/spf13/cobra"
)

// encrypterCmd represents the encrypter command
var encrypterCmd = &cobra.Command{
	Use:   "encrypter",
	Short: "A encrypter tool",
	Long: `A encrypter tool. For example:

加密：
./aphroditecli encrypter --secret "{SECRET}" "{RAW1}" "{RAW2}" "{RAW3}"

解密：
./aphroditecli encrypter -R --secret "{SECRET}" "{ENCRYPT1}" "{ENCRYPT2}" "{ENCRYPT3}"
`,
	Run: func(cmd *cobra.Command, args []string) {
		if reverse {
			encrypts.Decrypt(secret, args...)
		} else {
			encrypts.Encrypt(secret, args...)
		}
	},
}

func init() {
	rootCmd.AddCommand(encrypterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encrypterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encrypterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
