package utils

import (
	"duckdb-version-manager/models"
	"runtime"
)

type DeviceInfo struct {
	Architecture models.ArchitectureType
	Platform     models.PlatformType
}

func GetDeviceInfo() DeviceInfo {
	return DeviceInfo{
		Architecture: getArchitecture(),
		Platform:     getPlatform(),
	}
}

func getPlatform() models.PlatformType {
	os := runtime.GOOS
	if os == "linux" {
		return models.PlatformLinux
	} else if os == "darwin" {
		return models.PlatformMac
	} else if os == "windows" {
		return models.PlatformWindows
	}
	ExitWith("Unsupported platform: %s", os)
	return ""
}

func getArchitecture() models.ArchitectureType {
	arch := runtime.GOARCH
	if arch == "arm64" {
		return models.ArchitectureArm64
	} else if arch == "x86" {
		return models.ArchitectureX86
	}
	ExitWith("Unsupported architecture: %s", arch)
	return ""
}
