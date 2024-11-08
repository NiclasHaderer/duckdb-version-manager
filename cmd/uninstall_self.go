package cmd

import (
	"duckdb-version-manager/config"
	"duckdb-version-manager/utils"
	"github.com/spf13/cobra"
)

var uninstallSelfCmd = &cobra.Command{
	Use:   "uninstall-self",
	Short: "Removes duckman and all config files",
	Run: func(cmd *cobra.Command, args []string) {
		utils.RemoveFileOrDie(config.Dir)
		utils.RemoveFileOrDie(config.DuckmanBinaryFile)
		utils.RemoveFileOrDie(config.DefaultDuckdbFile)
	},
}
