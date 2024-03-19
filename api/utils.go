package api

import (
	"duckdb-version-manager/models"
	"duckdb-version-manager/utils"
)

func toVersionList(dict models.VersionDict) models.VersionList {
	versionList := make(models.VersionList, 0, len(dict))

	for versionStr, location := range dict {
		versionList = append(versionList, models.VersionInfo{Version: versionStr, RelativeVersionLocation: location})
	}
	utils.SortVersions(versionList, func(info models.VersionInfo) string {
		return info.Version
	})

	return versionList
}
