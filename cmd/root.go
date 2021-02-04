package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "os"
)

var rootCmd = &cobra.Command{
  Use:   "todo",
  Short: "todo",
  Long: `todo main command line tool`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Printf("args: %v", args)
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
