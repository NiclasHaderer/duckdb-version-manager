package models

type DuckdbReleases = map[string]DuckdbRelease

type DuckdbRelease struct {
	Version   string   `json:"version"`
	Name      string   `json:"name"`
	Platforms Platform `json:"platforms"`
}

type Platform = map[PlatformType]Architecture

type Architecture = map[ArchitectureType]ArchitectureRelease

type ArchitectureRelease struct {
	DownloadUrl string `json:"downloadUrl"`
}
