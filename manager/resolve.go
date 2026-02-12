package manager

import (
	"duckdb-version-manager/stacktrace"
	"fmt"
	"io"
	"strings"

	"github.com/hashicorp/go-version"
)

const motherDuckURL = "https://api.motherduck.com/latest_supported_duckdb_version.txt"

func (v *versionManagerImpl) resolveVersion(ver string) (string, stacktrace.Error) {
	switch strings.ToLower(ver) {
	case "md":
		return v.resolveMotherDuckVersion()
	case "latest":
		return v.resolveLatestVersion()
	default:
		return ver, nil
	}
}

func (v *versionManagerImpl) resolveMotherDuckVersion() (string, stacktrace.Error) {
	resp, err := v.client.Get().Get(motherDuckURL)
	if err != nil {
		return "", stacktrace.Wrap(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", stacktrace.Wrap(err)
	}

	resolved := strings.TrimSpace(string(body))
	fmt.Printf("Resolved 'md' to version %s\n", resolved)
	return resolved, nil
}

func (v *versionManagerImpl) resolveLatestVersion() (string, stacktrace.Error) {
	releases, sErr := v.client.ListAllReleasesDict()
	if sErr != nil {
		return "", sErr
	}

	var highest *version.Version
	var highestKey string

	for key := range releases {
		v, err := version.NewVersion(key)
		if err != nil {
			continue
		}
		if highest == nil || v.GreaterThan(highest) {
			highest = v
			highestKey = key
		}
	}

	if highest == nil {
		return "", stacktrace.NewF("Could not resolve 'latest' to a valid version")
	}

	fmt.Printf("Resolved 'latest' to version %s\n", highestKey)
	return highestKey, nil
}
