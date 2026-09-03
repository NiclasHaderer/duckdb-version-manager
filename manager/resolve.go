package manager

import (
	"duckdb-version-manager/models"
	"duckdb-version-manager/stacktrace"
	"fmt"
	"io"
	"strings"

	"github.com/hashicorp/go-version"
)

const motherDuckURL = "https://api.motherduck.com/latest_supported_duckdb_version.txt"
const stagingBaseURL = "https://duckdb-staging.duckdb.org"
const latestAlphaURL = stagingBaseURL + "/latest_alpha_version.txt"

// The alpha channel publishes one CLI tarball per platform/architecture pair.
var cliDistributions = map[models.PlatformType]map[models.ArchitectureType]string{
	models.PlatformWindows: {models.ArchitectureX86: "windows-amd64", models.ArchitectureArm64: "windows-arm64"},
	models.PlatformLinux:   {models.ArchitectureX86: "linux-amd64", models.ArchitectureArm64: "linux-arm64"},
	models.PlatformMac:     {models.ArchitectureX86: "osx-amd64", models.ArchitectureArm64: "osx-arm64"},
}

// resolveRelease turns a version or an alias into the release that should be installed.
// 'alpha' is the one alias that is not in the remote version index, because it points at a
// staged commit that moves too often to be tracked there.
func (v *versionManagerImpl) resolveRelease(ver string) (*models.VersionInformation, stacktrace.Error) {
	switch strings.ToLower(ver) {
	case "alpha":
		return v.resolveAlphaRelease()
	case "md":
		resolved, err := v.resolveMotherDuckVersion()
		if err != nil {
			return nil, err
		}
		ver = resolved
	case "latest":
		resolved, err := v.resolveLatestVersion()
		if err != nil {
			return nil, err
		}
		ver = resolved
	}

	return v.client.GetRelease(ver)
}

func cliRelease(ver string, baseURL string) *models.VersionInformation {
	platforms := make(models.Platforms, len(cliDistributions))
	for platform, architectures := range cliDistributions {
		platforms[platform] = make(models.Architectures, len(architectures))
		for architecture, distribution := range architectures {
			platforms[platform][architecture] = models.ArchitectureReleaseInformation{
				DownloadUrl: fmt.Sprintf("%s/duckdb-cli-%s.tar.gz", baseURL, distribution),
			}
		}
	}

	return &models.VersionInformation{Version: ver, Name: ver, Platforms: platforms}
}

// resolveAlphaRelease resolves the alpha channel to the staged commit it currently points at.
func (v *versionManagerImpl) resolveAlphaRelease() (*models.VersionInformation, stacktrace.Error) {
	body, err := v.getText(latestAlphaURL)
	if err != nil {
		return nil, err
	}

	commit, alphaVersion, found := strings.Cut(body, "/")
	if !found {
		return nil, stacktrace.NewF("Could not resolve 'alpha', unexpected version pointer '%s'", body)
	}
	if len(commit) > 10 {
		commit = commit[:10]
	}

	fmt.Printf("Resolved 'alpha' to version %s\n", alphaVersion)
	baseURL := fmt.Sprintf("%s/%s/%s/duckdb/duckdb/github_release", stagingBaseURL, commit, alphaVersion)
	return cliRelease(alphaVersion, baseURL), nil
}

func (v *versionManagerImpl) resolveMotherDuckVersion() (string, stacktrace.Error) {
	resolved, err := v.getText(motherDuckURL)
	if err != nil {
		return "", err
	}

	fmt.Printf("Resolved 'md' to version %s\n", resolved)
	return resolved, nil
}

func (v *versionManagerImpl) resolveLatestVersion() (string, stacktrace.Error) {
	releases, err := v.client.ListAllReleasesDict()
	if err != nil {
		return "", err
	}

	var highest *version.Version
	var highestKey string

	for key := range releases {
		parsed, parseErr := version.NewVersion(key)
		if parseErr != nil {
			continue
		}
		if highest == nil || parsed.GreaterThan(highest) {
			highest = parsed
			highestKey = key
		}
	}

	if highest == nil {
		return "", stacktrace.New("Could not resolve 'latest' to a valid version")
	}

	fmt.Printf("Resolved 'latest' to version %s\n", highestKey)
	return highestKey, nil
}

func (v *versionManagerImpl) getText(url string) (string, stacktrace.Error) {
	resp, err := v.client.Get().Get(url)
	if err != nil {
		return "", stacktrace.Wrap(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", stacktrace.Wrap(err)
	}

	return strings.TrimSpace(string(body)), nil
}
