/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"time"

	"github.com/illidaris/aphroditecli/internal/mongo"

	"github.com/spf13/cobra"
)

var (
	mongoconn   string
	mongodb     string
	concurrence int
)

// mongoCmd represents the mongo command
var mongoCmd = &cobra.Command{
	Use:   "mongo",
	Short: "A mongo tool",
	Long: `A mongo tool. For example:

./aphroditecli mongo --mongoconn "mongodb://{USER}:{PWD}@{IP}:27018/{DB}?authSource=admin" --mongodb "{DB}" --concurrence 10
}
	`,
	Run: func(cmd *cobra.Command, args []string) {
		//	mongo.ExecWithTrans(context.Background(), mongodb, mongoconn, concurrence)
		mongo.IndexSync(context.Background(), mongodb, mongoconn, concurrence)
		<-time.After(time.Hour * 1)
	},
}

func init() {
	rootCmd.AddCommand(mongoCmd)

	mongoCmd.PersistentFlags().StringVar(&mongoconn, "mongoconn", "", "mongo conn")
	mongoCmd.PersistentFlags().StringVar(&mongodb, "mongodb", "", "mongo db")
	mongoCmd.PersistentFlags().IntVar(&concurrence, "concurrence", 1, "concurrence")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mongoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mongoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
