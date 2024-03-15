package utils

import (
	"archive/zip"
	"bytes"
	"duckdb-version-manager/stacktrace"
	"io"
	"strings"
)

func getZipFileContent(file *zip.File) ([]byte, stacktrace.Error) {
	fileReader, err := file.Open()
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}
	defer fileReader.Close()

	content, err := io.ReadAll(fileReader)
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}
	return content, nil
}

func ExtractDuckdbFile(assetBytes []byte) ([]byte, stacktrace.Error) {
	zipReader, err := zip.NewReader(bytes.NewReader(assetBytes), int64(len(assetBytes)))
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}
	for _, file := range zipReader.File {
		fileName := file.Name

		// Normal release -> just one file with the name "duckdb"
		if fileName == "duckdb" {
			fileContent, err := getZipFileContent(file)
			if err != nil {
				return nil, err
			}

			return fileContent, nil
		} else if strings.Contains(fileName, "duckdb_cli") && strings.HasSuffix(fileName, ".zip") {
			// Nightly release -> multiple zip files within the zip file
			fileContent, err := getZipFileContent(file)
			if err != nil {
				return nil, err
			}
			return ExtractDuckdbFile(fileContent)
		}
	}

	return nil, stacktrace.New("Error during zip extraction: no file with the name 'duckdb' found in the zip file.")
}
