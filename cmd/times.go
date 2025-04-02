/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"time"

	"github.com/illidaris/aphrodite/pkg/convert"
	exptr "github.com/illidaris/aphroditecli/pkg/exporter"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

// timesCmd represents the times command
var timesCmd = &cobra.Command{
	Use:   "times",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		exptr.FmtTable([][]string{
			{"Date", "Ts", "DateFmt"},
			{convert.TimeFormat(now), cast.ToString(now.Unix()), cast.ToString(convert.TimeNumber(now))},
		}, false)
	},
}

func init() {
	rootCmd.AddCommand(timesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
