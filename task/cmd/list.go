package cmd

import (
	"fmt"
	"os"
	"stark/gophercises/task/db"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print the list of task",
	Long:  "This command is used to print all the tasks that a user has added",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()

		if err != nil {
			fmt.Println("Things went wrong: ", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks")
			return
		}

		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}
