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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
