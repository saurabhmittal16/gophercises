package cmd

import (
	"fmt"

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
		fmt.Println("This is a fake \"list\" command")
	},
}
