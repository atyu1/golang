/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [ID]",
	Short: "Complete task based on ID",
	Long: `Complete the line based on ID, search the list output to find the ID of item which needs to be completed,
	It will mark the Done field to true.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("Cannot properly get your ID from input: ", args, " Error: ", err)
		}

		initCSV(FILEPATH)
		line := findID(FILEPATH, uint32(id))
		if line == -1 {
			fmt.Println("No line for requested ID found, ID: ", args)
		}

		x := loadCSV(FILEPATH)
		x = deleteLine(x, line)
		replaceCSV(FILEPATH, x)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
