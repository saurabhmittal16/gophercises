package cmd

import (
	"errors"
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
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires the index of task")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		taskAdded := args[0]
		fmt.Printf("Added \"%s\" to your task list\n", taskAdded)
	},
}
