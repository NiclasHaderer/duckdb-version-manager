package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall [version]",
	Args:  cobra.ExactArgs(1),
	Short: "Uninstall a version of DuckDB",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("uninstall called")
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
