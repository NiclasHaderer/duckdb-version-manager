package utils

import (
	"duckdb-version-manager/models"
	"log"
	"strings"
)

func DownloadUrlTo(url string, dest string) error {
	println("Downloading", url, "to", dest)
	return nil
}

func GetDownloadUrlFrom(release *models.Release) string {
	sysInfo := GetDeviceInfo()

	platform, ok := release.Platforms[sysInfo.Platform]
	if !ok {
		log.Fatalf("Platform %s not supported. Supported architectures are: %s", sysInfo.Platform, strings.Join(Keys(release.Platforms), ", "))
	}

	arch, ok := platform[sysInfo.Architecture]
	if !ok {
		if res, ok := platform[models.ArchitectureUniversal]; ok {
			return res.DownloadUrl
		}

		log.Fatalf("Architecture %s not supported. Supported architectures are: %s", sysInfo.Architecture, strings.Join(Keys(platform), ", "))
	}

	return arch.DownloadUrl
}
