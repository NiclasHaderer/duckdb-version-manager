package cmd

import (
	"duckdb-version-manager/client"
	"duckdb-version-manager/config"
	"duckdb-version-manager/utils"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [version]",
	Short: "Install a specific version of DuckDB",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiClient := client.New()
		version := args[0]

		possibleVersions, err := apiClient.ListAllReleasesDict()
		if err != nil {
			utils.ExitWithError(err)
		}
		versionLocation, ok := possibleVersions[version]
		if !ok {
			utils.ExitWith("Version '%s' not found", version)
		}

		resolvedVersion, err := apiClient.GetReleaseWithLocation(versionLocation)
		if err != nil {
			utils.ExitWithError(err)
		}

		downloadUrl := utils.GetDownloadUrlFrom(resolvedVersion)
		err = utils.DownloadUrlTo(downloadUrl, config.InstallDir+"/"+resolvedVersion.Version)
		if err != nil {
			utils.ExitWithError(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
