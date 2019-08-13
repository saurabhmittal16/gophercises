package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark task as done",
	Long:  "This command is used to mark a task in your list of tasks as done",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires the index of task")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		taskInd, err := strconv.Atoi(args[0])
		if err == nil {
			fmt.Printf("You have completed task %d\n", taskInd)
		} else {
			fmt.Println("Invalid arguments")
		}
	},
}
