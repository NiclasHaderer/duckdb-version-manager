package models

type VersionStr = string

// Models that describe a specific version of DuckDB

type VersionInformation struct {
	Version   VersionStr `json:"version"`
	Name      string     `json:"name"`
	Platforms Platforms  `json:"platforms"`
}

type Platforms = map[PlatformType]Architectures

type Architectures = map[ArchitectureType]ArchitectureReleaseInformation

type ArchitectureReleaseInformation struct {
	DownloadUrl string `json:"downloadUrl"`
}

// Models that describe the overview of the remote versions, not the remote versions themselves

type RemoteVersionPath = string
type RemoteVersionDict = map[VersionStr]RemoteVersionPath

type RemoteVersionInfo struct {
	Version                 VersionStr
	RelativeVersionLocation RemoteVersionPath
}
type RemoteVersions = []RemoteVersionInfo
