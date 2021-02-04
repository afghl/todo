package cmd

import (
  "fmt"
  "strings"
  "github.com/spf13/cobra"
)

var echoCmd = &cobra.Command{
  Use:   "echo",
  Short: "echo what was input",
  Long:  `echo`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Echo: " + strings.Join(args, " "))
  },
}

func init() {
	rootCmd.AddCommand(echoCmd)
}
