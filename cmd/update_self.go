package cmd

import (
	"duckdb-version-manager/api"
	"duckdb-version-manager/config"
	"duckdb-version-manager/utils"
	"github.com/spf13/cobra"
	"time"
)

var updateSelfCmd = &cobra.Command{
	Use:   "update-self",
	Short: "Updates duckman to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		client := api.New()
		release, err := client.LatestDuckVmRelease(time.Minute * 5)
		if err != nil {
			utils.ExitWithError(err)
		}

		downloadUrl, err := utils.GetDownloadUrlFrom(release)
		if err != nil {
			utils.ExitWithError(err)
		}

		body, err := utils.GetResponseBodyFrom(client.Get(), *downloadUrl)
		if err != nil {
			utils.ExitWithError(err)
		}

		if err := utils.WriteFileAtomic(config.DuckmanBinaryFile, body, 0700); err != nil {
			utils.ExitWithError(err)
		}
	},
}
