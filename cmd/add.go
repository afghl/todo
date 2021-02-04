package cmd

import (
  "github.com/spf13/cobra"
  "strings"
  "todo/persist"
  "todo/todos"
)

var addCmd = &cobra.Command{
  Use:   "add",
  Short: "Add a todo todos.",
  Run: func(cmd *cobra.Command, args []string) {
    content := strings.Join(args, " ")
    todo := todos.NewTodo(content)
    // load existed items
    todos, _ := persist.Load()
    todos.Append(*todo)
    persist.Save(todos)
  },
}

func init() {
	rootCmd.AddCommand(addCmd)
}
