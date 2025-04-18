/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"time"

	"github.com/illidaris/aphroditecli/internal/database"
	"github.com/spf13/cobra"
)

var (
	trans int32
)

// dbexecCmd represents the dbexec command
var dbexecCmd = &cobra.Command{
	Use:   "dbexec",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		database.DbExec(dbDsn, trans, time.Duration(delay)*time.Second, args...)
	},
}

func init() {
	rootCmd.AddCommand(dbexecCmd)
	dbexecCmd.PersistentFlags().StringVar(&dbDsn, "dbDsn", "", "db dsn")
	dbexecCmd.PersistentFlags().Int32Var(&trans, "trans", -1, "trans level: -1 no trans, 0 default, 1 LevelReadUncommitted, 2 LevelReadCommitted, 3 LevelWriteCommitted, 4 LevelRepeatableRead, 5 LevelSnapshot, 6 LevelSerializable, 7 LevelLinearizable")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbexecCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbexecCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
