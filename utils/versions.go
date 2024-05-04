package utils

import (
	"duckdb-version-manager/models"
	"github.com/hashicorp/go-version"
	"sort"
)

func SortVersions[T any](versions []T, getVersion func(T) string) {
	sort.Slice(versions, func(i, j int) bool {
		if getVersion(versions[i]) == "nightly" {
			return true
		}
		if getVersion(versions[j]) == "nightly" {
			return false

		}
		v1, _ := version.NewVersion(getVersion(versions[i]))
		v2, _ := version.NewVersion(getVersion(versions[j]))

		return v2.LessThan(v1)
	})
}

func ToVersionList(dict *models.RemoteVersionDict) models.RemoteVersions {
	versionList := make(models.RemoteVersions, 0, len(*dict))

	for versionStr, location := range *dict {
		versionList = append(versionList, models.RemoteVersionInfo{Version: versionStr, RelativeVersionLocation: location})
	}
	SortVersions(versionList, func(info models.RemoteVersionInfo) string {
		return info.Version
	})

	return versionList
}
