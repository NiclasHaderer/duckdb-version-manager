package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available DuckDB versions. Use 'local' to list local versions and 'remote' to list remote versions.",
}
