package list

import (
	"github.com/spf13/cobra"
)

var LocalCmd = &cobra.Command{
	Use:   "local",
	Short: "List locally installed DuckDB versions",
	Run: func(cmd *cobra.Command, args []string) {
		// List the locally installed DuckDB versions

	},
}
