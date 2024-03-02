package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// defaultCmd represents the default command
var defaultCmd = &cobra.Command{
	Use:   "default [version]",
	Short: "Set an already installed version of DuckDB as the default one",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("default called")
	},
}

func init() {
	rootCmd.AddCommand(defaultCmd)
}
