/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/illidaris/aphroditecli/internal/qrcodes"
	"github.com/spf13/cobra"
)

var (
	logo  string
	logoP int
	zoom  int
)

// qrcodeCmd represents the qrcode command
var qrcodeCmd = &cobra.Command{
	Use:   "qrcode",
	Short: "A brief description of your command, default is encode generate a qrcode png file, -R decode read a qrcode png file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

default is encode generate a qrcode png file, -R decode read a qrcode png file

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if reverse {
			qrcodes.ParseQrCodeExport(args...)
		} else {
			qrcodes.WriteQrCodeExport(size, logoP, zoom, logo, out, args...)
		}
	},
}

func init() {
	rootCmd.AddCommand(qrcodeCmd)
	qrcodeCmd.PersistentFlags().IntVar(&size, "size", 256, "qrcode size")
	qrcodeCmd.PersistentFlags().StringVar(&logo, "logo", "", "qrcode with logo")
	qrcodeCmd.PersistentFlags().IntVar(&logoP, "logoP", 5, "qrcode proportion [0,10] default 5")
	qrcodeCmd.PersistentFlags().IntVar(&zoom, "zoom", 9, "qrcode zoom [0,10] default 9")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// qrcodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// qrcodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
