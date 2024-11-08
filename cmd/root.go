package cmd

import (
	"duckdb-version-manager/models"
	"duckdb-version-manager/utils"
	"github.com/spf13/cobra"
	"os"
)

func getCommandUse() string {
	deviceInfo := utils.GetDeviceInfo()
	baseCmd := "duckman"
	if deviceInfo.Platform == models.PlatformWindows {
		baseCmd += ".exe"
	}
	return baseCmd
}

var rootCmd = &cobra.Command{
	Use:   getCommandUse(),
	Short: "A version manager for DuckDB",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
