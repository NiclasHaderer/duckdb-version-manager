package models

type PlatformType = int

const (
	PlatformWindows PlatformType = iota
	PlatformMac     PlatformType = iota
	PlatformLinux   PlatformType = iota
)
