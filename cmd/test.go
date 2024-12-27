/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"loadtest/lib"
	"log"

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

			fmt.Println("Testing url: ", url)
			fmt.Println("Number of requests: ", n)
			lib.TestUrl(url, n, c)
			defer log.Println("Test completed")
		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	testCmd.Flags().IntP("requests", "n", 1, "Number of requests")

	testCmd.Flags().IntP("concurrency", "c", 1, "Number of concurrent requests")

}
