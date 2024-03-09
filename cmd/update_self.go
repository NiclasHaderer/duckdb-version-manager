package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var updateSelfCmd = &cobra.Command{
	Use:   "update-self",
	Short: "Updates duck-vm to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatalf("This feature is not implemented yet")
	},
}

func init() {
	rootCmd.AddCommand(updateSelfCmd)
}
