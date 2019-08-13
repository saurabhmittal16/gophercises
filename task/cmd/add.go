package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",
	Long:  "This command is used to add a task to your list of tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			taskAdded := args[0]
			fmt.Printf("Added \"%s\" to your task list\n", taskAdded)
		} else {
			fmt.Println("Command add requires the task to be added")
		}
	},
}
