package cmd

import (
	"duckdb-version-manager/cmd/list"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available DuckDB versions",
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(list.LocalCmd)
	listCmd.AddCommand(list.RemoteCmd)
}
