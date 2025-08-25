/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	"github.com/illidaris/aphroditecli/internal/database"
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
	Short: "A db query tool",
	Long: `A db query tool. For example:

./aphroditecli db --dbDsn '{USER}:{PWD}@tcp({IP}:3306)/{DB}' --dbDriver mysql --dbSql "select * from ..." --out excel --pretty
`,
	Run: func(cmd *cobra.Command, args []string) {
		err := database.DbQueryExport(context.Background(),
			args, out, pretty,
			database.WithDriver(dbDriver),
			database.WithDSN(dbDsn),
		)
		if err != nil {
			println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.PersistentFlags().StringVar(&dbDriver, "dbDriver", "mysql", "db driver")
	dbCmd.PersistentFlags().StringVar(&dbDsn, "dbDsn", "", "db dsn")
	dbCmd.PersistentFlags().StringVar(&dbSql, "dbSql", "", "db sql")
}
