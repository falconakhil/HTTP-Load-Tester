/*
Copyright © 2024 AKHIL SAKTHIESWARAN <EMAIL ADDRESS>
*/
package cmd

import (
	"loadtest/lib"

	"github.com/spf13/cobra"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Test urls from a csv file",
	Long:  `Test urls from a csv file`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			filePath := args[0]
			lib.TestFile(filePath)

		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
