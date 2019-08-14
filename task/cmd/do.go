package cmd

import (
	"errors"
	"fmt"
	"os"
	"stark/gophercises/task/db"
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
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse argument: ", arg)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Things went wrong: ", err)
			os.Exit(1)
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}

			task := tasks[id-1]
			err := db.DeleteTask(task.Key)

			if err != nil {
				fmt.Printf("Failed to mark \"%s\" as complete. Error: %s\n", task.Value, err)
			} else {
				fmt.Printf("Marked \"%s\" as complete\n", task.Value)
			}
		}
	},
}
