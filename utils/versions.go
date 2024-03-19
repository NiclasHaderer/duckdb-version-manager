package utils

import (
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
