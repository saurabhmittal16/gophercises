package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark task as done",
	Long:  "This command is used to mark a task in your list of tasks as done",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			taskInd := args[0]
			fmt.Printf("You have completed the \"%s\" task.\n", taskInd)
		} else {
			fmt.Println("Command do requires the task index")
		}
	},
}
