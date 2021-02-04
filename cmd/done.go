package cmd

import (
	"github.com/spf13/cobra"
	"strconv"
	"todo/persist"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Set a todo done.",
	Run: func(cmd *cobra.Command, args []string) {
		idStr := args[1]
		// load existed items
		todos, _ := persist.Load()
		id, _ := strconv.Atoi(idStr)
		todo := todos.Find(id)
		if todo != nil {
			todo.Done()
		}
		persist.Save(todos)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}