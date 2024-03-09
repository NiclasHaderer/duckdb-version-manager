package cmd

import (
	"duckdb-version-manager/client"
	"duckdb-version-manager/config"
	"duckdb-version-manager/utils"
	"github.com/spf13/cobra"
)

var updateSelfCmd = &cobra.Command{
	Use:   "update-self",
	Short: "Updates duck-vm to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		client := client.New()
		release, err := client.LatestDuckVmRelease()
		if err != nil {
			utils.ExitWithError(err)
		}

		downloadUrl, err := utils.GetDownloadUrlFrom(release)
		if err != nil {
			utils.ExitWithError(err)
		}

		err = utils.DownloadUrlTo(downloadUrl, config.InstallDir+"/"+config.DuckVMName, false)
		if err != nil {
			utils.ExitWithError(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateSelfCmd)
}
