package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var updateSelfCmd = &cobra.Command{
	Use:   "updateSelf",
	Short: "Updates the duck-vm binary to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatalf("This feature is not implemented yet")
	},
}

func init() {
	rootCmd.AddCommand(updateSelfCmd)
}
