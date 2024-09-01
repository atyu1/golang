/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [ToDo Description]",
	Short: "Add a new ToDo Item to the list",
	Long:  `Adds a new ToDo Item to the list, please specify description. By default it will add actual time in UTC and also mark as not done yet.`,
	Run: func(cmd *cobra.Command, args []string) {
		initCSV(FILEPATH)

		lastId := getLastID(FILEPATH)
		nextId := lastId + 1

		//item := []string{nextId, "Shopping with my best friend and his best friends of the friends", "Tomorrow", "false"}
		item := todoItems{
			ID:          nextId,
			Description: strings.Join(args, " "),
			DateCreated: time.Now().UTC(),
			IsCompleted: false,
		}

		appendCSV(FILEPATH, item)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
