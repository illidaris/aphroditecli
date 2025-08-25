/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	"github.com/illidaris/aphroditecli/internal/ollama"
	"github.com/spf13/cobra"
)

var (
	ollamaAction       string
	ollamaHost         string
	ollamaModel        string
	ollamaTemplate     string
	ollamaLabelFile    string
	ollamaCategoryFile string
)

// ollamaCmd represents the ollama command
var ollamaCmd = &cobra.Command{
	Use:   "ollama",
	Short: "A ollama tool",
	Long: `A ollama tool. For example:

./aphroditecli.exe ollama --ollamaAction classic --ollamaHost "{OLLAMA_HOST}" --ollamaLabelFile "{LABEL_PATH}" --ollamaCategoryFile "{CATEGORY_PATH}" --out excel "{RAW_PATH}"
`,
	Run: func(cmd *cobra.Command, args []string) {
		switch ollamaAction {
		case "classic":
			err := ollama.Classic(context.Background(), ollamaHost, ollamaModel, ollamaTemplate, ollamaLabelFile, ollamaCategoryFile, out, args)
			if err != nil {
				println(err.Error())
			}
		default:
			println("unknown action")
		}
	},
}

func init() {
	rootCmd.AddCommand(ollamaCmd)

	ollamaCmd.PersistentFlags().StringVar(&ollamaAction, "ollamaAction", "classic", "ollama action")
	ollamaCmd.PersistentFlags().StringVar(&ollamaHost, "ollamaHost", "http://localhost:11434", "ollama host, default http://localhost:11434")
	ollamaCmd.PersistentFlags().StringVar(&ollamaModel, "ollamaModel", "deepseek-r1:32b", "ollama model, default is deepseek-r1:32b")
	ollamaCmd.PersistentFlags().StringVar(&ollamaTemplate, "ollamaTemplate", "", "ollama template")
	ollamaCmd.PersistentFlags().StringVar(&ollamaLabelFile, "ollamaLabelFile", "./label.xlsx", "ollama label file, default is label.xlsx")
	ollamaCmd.PersistentFlags().StringVar(&ollamaCategoryFile, "ollamaCategoryFile", "./category.xlsx", "ollama label file, default is category.xlsx")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ollamaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ollamaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
