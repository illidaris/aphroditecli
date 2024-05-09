/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"aphroditecli/internal/watermill"
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	kafkaAddrs  []string
	kafkaUser   string
	kafkaPwd    string
	kafkaTopics []string
	kafkaMode   string
	kafkaKey    string
	kafkaValue  string
)

// watermillCmd represents the watermill command
var watermillCmd = &cobra.Command{
	Use:   "watermill",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("watermill called")
		fmt.Println(strings.Join(kafkaAddrs, ","))
		fmt.Println("kafka called")
		if kafkaMode == "producer" {
			if len(kafkaTopics) == 0 {
				println("topic is nil")
				return
			}
			err := watermill.Publish(
				context.Background(),
				kafkaAddrs,
				kafkaUser,
				kafkaPwd,
				kafkaTopics[0],
				kafkaKey,
				kafkaValue)
			if err != nil {
				println(err.Error())
			}
		} else if kafkaMode == "consumer" {
			err := watermill.Consumer(
				context.Background(),
				kafkaTopics,
				kafkaAddrs,
				kafkaUser,
				kafkaPwd,
				delay,
			)
			if err != nil {
				println(err.Error())
			} else {
				println("listen: ", strings.Join(kafkaTopics, ","), "...")
				ch := make(chan struct{})
				<-ch
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(watermillCmd)

	watermillCmd.PersistentFlags().StringSliceVarP(&kafkaAddrs, "addrs", "A", []string{}, "kafka address")
	watermillCmd.PersistentFlags().StringVarP(&kafkaUser, "user", "U", "", "kafka user")
	watermillCmd.PersistentFlags().StringVarP(&kafkaPwd, "pwd", "P", "", "kafka pwd")
	watermillCmd.PersistentFlags().StringSliceVarP(&kafkaTopics, "topics", "T", []string{}, "kafka topics")
	watermillCmd.PersistentFlags().StringVarP(&kafkaMode, "mode", "M", "", "kafka mode producer or consumer")
	watermillCmd.PersistentFlags().StringVarP(&kafkaKey, "key", "K", "", "kafka message key")
	watermillCmd.PersistentFlags().StringVarP(&kafkaValue, "value", "V", "", "kafka message value")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// watermillCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// watermillCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
