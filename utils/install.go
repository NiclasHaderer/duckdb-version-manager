package utils

import (
	"duckdb-version-manager/client"
	"duckdb-version-manager/config"
	"fmt"
)

func InstallVersion(version string) error {
	apiClient := client.New()

	possibleVersions, err := apiClient.ListAllReleasesDict()
	if err != nil {
		return err
	}
	versionLocation, ok := possibleVersions[version]
	if !ok {
		return fmt.Errorf("version '%s' not found", version)
	}

	resolvedVersion, err := apiClient.GetReleaseWithLocation(versionLocation)
	if err != nil {
		return err
	}

	downloadUrl, err := GetDownloadUrlFrom(resolvedVersion)
	if err != nil {
		return err
	}
	err = DownloadUrlTo(downloadUrl, config.VersionDir+"/"+resolvedVersion.Version)
	if err != nil {
		return err
	}

	return nil
}
