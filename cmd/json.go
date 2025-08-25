/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "A json pretty tool",
	Long: `A json pretty tool. For example:
./aphroditecli json --pretty "{JSON DATA}"
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("json called")
		for _, arg := range args {
			println(PrettyString(arg))
		}
	},
}

func PrettyString(str string) string {
	if !pretty {
		return str
	}
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return str
	}
	return prettyJSON.String()
}

func init() {
	rootCmd.AddCommand(jsonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jsonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
