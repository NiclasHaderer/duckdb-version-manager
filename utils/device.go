package utils

import (
	"duckdb-version-manager/models"
	"duckdb-version-manager/stacktrace"
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

//goland:noinspection GoBoolExpressions
func getPlatform() models.PlatformType {
	os := runtime.GOOS
	if os == "linux" {
		return models.PlatformLinux
	} else if os == "darwin" {
		return models.PlatformMac
	} else if os == "windows" {
		return models.PlatformWindows
	}
	err := stacktrace.NewF("Unsupported platform: %s", os)
	ExitWithError(err)
	return ""
}

//goland:noinspection GoBoolExpressions
func getArchitecture() models.ArchitectureType {
	arch := runtime.GOARCH
	if arch == "arm64" {
		return models.ArchitectureArm64
	} else if arch == "x86" || arch == "amd64" {
		return models.ArchitectureX86
	}
	err := stacktrace.NewF("Unsupported architecture: %s", arch)
	ExitWithError(err)
	return ""
}
