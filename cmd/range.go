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
var rangeCmd = &cobra.Command{
	Use:   "range",
	Short: "Load test a given url for a range of concurrent users",
	Long:  `Load test a given url for a range of concurrent users and generate graphs for it`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) >= 1 {
			url := args[0]
			cBegin, _ := cmd.Flags().GetInt("concurrency_start")
			cEnd,_ :=cmd.Flags().GetInt("concurrency_end")
			cStep,_:=cmd.Flags().GetInt("concurrency_step")
			output_dir,_:=cmd.Flags().GetString("output")
			n, _ := cmd.Flags().GetInt("requests")

			// Set the log file path
			logFilePath, _ := cmd.Flags().GetString("logfile")
			fmt.Println("Log file path: ", logFilePath)
			logFile, _ := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			log.SetOutput(logFile)

			lib.TestRange(url, n, cBegin,cEnd,cStep,output_dir)
			defer log.Println("Test completed")
		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(rangeCmd)

	rangeCmd.Flags().IntP("requests", "n", 1, "Number of requests")
	rangeCmd.Flags().IntP("concurrency_start", "s", 1, "Beginning of concurrency range")
	rangeCmd.Flags().IntP("concurrency_end", "e", 2, "End of concurrency range")
	rangeCmd.Flags().IntP("concurrency_step", "j", 1, "Step for concurrency range")
	rangeCmd.Flags().StringP("output","o","./output/","Ouptut dir")
}
