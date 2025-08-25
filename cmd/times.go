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

var (
	isTs bool
)

// timesCmd represents the times command
var timesCmd = &cobra.Command{
	Use:   "times",
	Short: "A unix tool",
	Long: `A unix tool. For example:

./aphroditecli times --isTs "{TS1}" "{TS2}"
./aphroditecli times "{DATE1}" "{DATE2}"
`,
	Run: func(cmd *cobra.Command, args []string) {
		ts := []time.Time{}
		for _, v := range args {
			if isTs {
				ts = append(ts, time.Unix(cast.ToInt64(v), 0))
			} else {
				ts = append(ts, cast.ToTime(v))
			}
		}
		if len(ts) == 0 {
			ts = append(ts, time.Now())
		}
		rows := [][]string{
			{"Date", "Ts", "DateFmt", "DateFmtCN"},
		}
		for _, v := range ts {
			row := []string{convert.TimeFormat(v), cast.ToString(v.Unix()), cast.ToString(convert.TimeNumber(v)), v.Format("2006年01月02日 15:04:05")}
			rows = append(rows, row)
		}
		exptr.FmtTable(rows, false)
	},
}

func init() {
	rootCmd.AddCommand(timesCmd)
	timesCmd.PersistentFlags().BoolVar(&isTs, "isTs", false, "input args is timestamp")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
