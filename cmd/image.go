/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/illidaris/aphroditecli/internal/images"
	"github.com/spf13/cobra"
)

var (
	width  uint
	height uint
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, v := range args {
			err := images.Resize(v, out, width, height)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)

	imageCmd.PersistentFlags().UintVar(&width, "width", 100, "width size")
	imageCmd.PersistentFlags().UintVar(&height, "height", 100, "height size")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// imageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// imageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
