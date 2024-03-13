package utils

import (
	"duckdb-version-manager/models"
	"io"
	"log"
	"net/http"
	"strings"
)

func GetResponseBodyFrom(client *http.Client, url string) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func GetDownloadUrlFrom(release *models.Release) (string, error) {
	sysInfo := GetDeviceInfo()

	platform, ok := release.Platforms[sysInfo.Platform]
	if !ok {
		log.Fatalf("Platform %s not supported. Supported architectures are: %s", sysInfo.Platform, strings.Join(Keys(release.Platforms), ", "))
	}

	arch, ok := platform[sysInfo.Architecture]
	if !ok {
		if res, ok := platform[models.ArchitectureUniversal]; ok {
			return res.DownloadUrl, nil
		}

		log.Fatalf("Architecture %s not supported. Supported architectures are: %s", sysInfo.Architecture, strings.Join(Keys(platform), ", "))
	}

	return arch.DownloadUrl, nil
}
