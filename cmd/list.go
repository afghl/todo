package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"todo/persist"
	"todo/todos"
)

var list = &cobra.Command{
	Use:   "list",
	Short: "List all todos.",
	Run: func(cmd *cobra.Command, args []string) {
		queryAll := len(args) == 1 && args[0] == "all"
		// load existed items
		todoItems, _ := persist.Load()

		printTodo(todoItems, queryAll)
	},
}

func printTodo(todoItems *todos.Todos, all bool) {
	doneCnt := 0
	todoItems.Iterate(func(i int, t todos.Todo) {
		if !all && t.IsDone() {
			return
		}
		doneCnt++
		var done string
		if t.IsDone() {
			done = "[Done]"
		} else {
			done = ""
		}
		fmt.Printf(strconv.Itoa(t.Id) + ". " + done + " " + t.Content + ". \n")
	})
	fmt.Printf("Total: %v, %v item done \n", todoItems.Count(), doneCnt)
}

func init() {
	rootCmd.AddCommand(list)
}
