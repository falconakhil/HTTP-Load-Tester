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

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Load test a given url",
	Long:  `Load test a given url`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) >= 1 {
			url := args[0]
			c, _ := cmd.Flags().GetInt("concurrency")
			n, _ := cmd.Flags().GetInt("requests")

			// Set the log file path
			logFilePath, _ := cmd.Flags().GetString("logfile")
			fmt.Println("Log file path: ", logFilePath)
			logFile, _ := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			log.SetOutput(logFile)

			lib.TestUrl(url, n, c)
			defer log.Println("Test completed")
		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	testCmd.Flags().IntP("requests", "n", 1, "Number of requests")
	testCmd.Flags().IntP("concurrency", "c", 1, "Number of concurrent requests")
}
