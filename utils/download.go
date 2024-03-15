package utils

import (
	"duckdb-version-manager/models"
	"duckdb-version-manager/stacktrace"
	"io"
	"log"
	"net/http"
	"strings"
)

func GetResponseBodyFrom(client *http.Client, url string) ([]byte, stacktrace.Error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}
	return body, nil
}

func GetDownloadUrlFrom(release *models.Release) (*string, stacktrace.Error) {
	sysInfo := GetDeviceInfo()

	platform, ok := release.Platforms[sysInfo.Platform]
	if !ok {
		log.Fatalf("Platform %s not supported. Supported architectures are: %s", sysInfo.Platform, strings.Join(Keys(release.Platforms), ", "))
	}

	arch, ok := platform[sysInfo.Architecture]
	if !ok {
		if res, ok := platform[models.ArchitectureUniversal]; ok {
			return &res.DownloadUrl, nil
		}

		return nil, stacktrace.NewF("Architecture %s not supported. Supported architectures are: %s", sysInfo.Architecture, strings.Join(Keys(platform), ", "))
	}

	return &arch.DownloadUrl, nil
}
