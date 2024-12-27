/*
Copyright © 2024 AKHIL SAKTHIESWARAN <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"loadtest/lib"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Test urls from a csv file",
	Long:  `Test urls from a csv file`,
	Run: func(cmd *cobra.Command, args []string) {

		c, _ := cmd.Flags().GetInt("concurrent_urls")

		// Set the log file path
		logFilePath, _ := cmd.Flags().GetString("logfile")
		fmt.Println("Log file path: ", logFilePath)
		logFile, _ := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		log.SetOutput(logFile)


		if len(args) >= 1 {
			filePath := args[0]
			lib.TestFile(filePath, c)

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

	fileCmd.Flags().IntP("concurrent_urls", "c", 1, "Number of concurrent urls")
}
