/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	"github.com/illidaris/aphroditecli/internal/database"
	exptr "github.com/illidaris/aphroditecli/pkg/exporter"
	"github.com/spf13/cobra"
)

var (
	dbDsn    string
	dbSql    string
	dbDriver string
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := database.DbExec(context.Background(),
			dbSql, nil,
			database.WithDriver(dbDriver),
			database.WithDSN(dbDsn),
		)
		if err != nil {
			println(err.Error())
		}
		exptr.Export(out, data, pretty)
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.PersistentFlags().StringVar(&dbDriver, "dbDriver", "mysql", "db driver")
	dbCmd.PersistentFlags().StringVar(&dbDsn, "dbDsn", "", "db dsn")
	dbCmd.PersistentFlags().StringVar(&dbSql, "dbSql", "", "db sql")
}
