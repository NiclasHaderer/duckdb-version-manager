package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use [version]",
	Short: "Use an already installed version of DuckDB for this shell session",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("use called")
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
