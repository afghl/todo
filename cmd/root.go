package cmd

import (
  "os"
  "fmt"
  "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "todo",
  Short: "todo",
  Long: `todo main command line tool`,
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
