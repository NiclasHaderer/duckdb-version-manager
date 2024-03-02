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

type Version = string
type RelativeVersionLocation = string
type VersionDict = map[Version]RelativeVersionLocation

// List of tuples

type VersionInfo struct {
	Version                 Version
	RelativeVersionLocation RelativeVersionLocation
}
type VersionList = []VersionInfo
