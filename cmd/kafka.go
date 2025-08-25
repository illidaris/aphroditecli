/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/illidaris/aphroditecli/internal/kafka"

	"github.com/illidaris/aphrodite/component/kafkaex"
	"github.com/spf13/cobra"
)

// kafkaCmd represents the kafka command
var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "A kafka tool",
	Long: `A kafka tool. For example:
consumer:
./aphroditecli kafka -A192.168.97.71:10192 -A192.168.97.71:10292 -A192.168.97.71:10392 -Ukafka -PpUuQNY9zG3NvObZxDwBhiHSBD6UxsQVx -Ttestdelaytopic -Mconsumer

producer:
./aphroditecli kafka -A192.168.97.71:10192 -A192.168.97.71:10292 -A192.168.97.71:10392 -Ukafka -PpUuQNY9zG3NvObZxDwBhiHSBD6UxsQVx -Ttestdelaytopic -Mproducer -Kabcd -Vmessageaaaaaaaaaaaaaaaaaaaaaaaaaaa
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strings.Join(kafkaAddrs, ","))
		fmt.Println("kafka called")
		opts := []kafkaex.OptionsFunc{
			kafkaex.WithAddr(kafkaAddrs...),
			kafkaex.WithApp("github.com/illidaris/aphroditecli"),
			kafkaex.WithUser(kafkaUser),
			kafkaex.WithPwd(kafkaPwd),
		}
		if kafkaMode == "producer" {
			if len(kafkaTopics) == 0 {
				println("topic is nil")
				return
			}
			err := kafka.Publish(context.Background(), kafkaTopics[0], kafkaKey, kafkaValue, opts...)
			if err != nil {
				println(err.Error())
			}
		} else if kafkaMode == "consumer" {
			err := kafka.Consumer(context.Background(), kafkaTopics, opts...)
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
	rootCmd.AddCommand(kafkaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	kafkaCmd.PersistentFlags().StringSliceVarP(&kafkaAddrs, "addrs", "A", []string{}, "kafka address")
	kafkaCmd.PersistentFlags().StringVarP(&kafkaUser, "user", "U", "", "kafka user")
	kafkaCmd.PersistentFlags().StringVarP(&kafkaPwd, "pwd", "P", "", "kafka pwd")
	kafkaCmd.PersistentFlags().StringSliceVarP(&kafkaTopics, "topics", "T", []string{}, "kafka topics")
	kafkaCmd.PersistentFlags().StringVarP(&kafkaMode, "mode", "M", "", "kafka mode producer or consumer")
	kafkaCmd.PersistentFlags().StringVarP(&kafkaKey, "key", "K", "", "kafka message key")
	kafkaCmd.PersistentFlags().StringVarP(&kafkaValue, "value", "V", "", "kafka message value")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kafkaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
