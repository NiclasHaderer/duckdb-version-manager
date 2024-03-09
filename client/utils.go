package client

import (
	"duckdb-version-manager/models"
	"github.com/hashicorp/go-version"
	"sort"
)

func toVersionList(dict models.VersionDict) models.VersionList {
	versionList := make(models.VersionList, 0, len(dict))

	for versionStr, location := range dict {
		versionList = append(versionList, models.VersionInfo{Version: versionStr, RelativeVersionLocation: location})
	}

	// Sort the version list
	sort.Slice(versionList, func(i, j int) bool {
		if versionList[i].Version == "nightly" {
			return true
		}
		if versionList[j].Version == "nightly" {
			return false

		}
		v1, _ := version.NewVersion(versionList[i].Version)
		v2, _ := version.NewVersion(versionList[j].Version)

		return v2.LessThan(v1)
	})
	return versionList
}
