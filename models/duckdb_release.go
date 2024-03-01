package models

type Releases = map[string]Release

type Release struct {
	Version   string    `json:"version"`
	Name      string    `json:"name"`
	Platforms Platforms `json:"platforms"`
}

type Platforms = map[PlatformType]Architectures

type Architectures = map[ArchitectureType]ArchitectureReleaseInformation

type ArchitectureReleaseInformation struct {
	DownloadUrl string `json:"downloadUrl"`
}

type VersionName = string
type RelativeVersionLocation = string
type VersionDict = map[VersionName]RelativeVersionLocation

// List of tuples

type VersionInfo struct {
	VersionName             VersionName
	RelativeVersionLocation RelativeVersionLocation
}
type VersionList = []VersionInfo
