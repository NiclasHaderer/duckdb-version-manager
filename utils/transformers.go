package utils

import (
	"duckdb-version-manager/models"
	"github.com/hashicorp/go-version"
	"sort"
)

func ToVersionList(dict models.VersionDict) models.VersionList {
	versionList := make(models.VersionList, 0, len(dict))

	for versionStr, location := range dict {
		versionList = append(versionList, models.VersionInfo{Version: versionStr, RelativeVersionLocation: location})
	}

	// Sort the version list
	sort.Slice(versionList, func(i, j int) bool {
		v1, _ := version.NewVersion(versionList[i].Version)
		v2, _ := version.NewVersion(versionList[j].Version)

		return v2.LessThan(v1)
	})
	return versionList
}
