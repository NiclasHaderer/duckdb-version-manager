package utils

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"duckdb-version-manager/stacktrace"
	"io"
	"strings"
)

func isDuckdbBinary(name string) bool {
	name = name[strings.LastIndexAny(name, "/\\")+1:]
	return name == "duckdb" || name == "duckdb.exe"
}

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

// ExtractDuckdbFile pulls the duckdb binary out of a release asset, which is either a
// gzipped tarball or a zip file (which in turn may contain the actual per-platform zip).
func ExtractDuckdbFile(assetBytes []byte) ([]byte, stacktrace.Error) {
	if bytes.HasPrefix(assetBytes, []byte{0x1f, 0x8b}) {
		return extractFromTarGz(assetBytes)
	}
	return extractFromZip(assetBytes)
}

func extractFromTarGz(assetBytes []byte) ([]byte, stacktrace.Error) {
	gzipReader, err := gzip.NewReader(bytes.NewReader(assetBytes))
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, stacktrace.Wrap(err)
		}

		if header.Typeflag != tar.TypeReg || !isDuckdbBinary(header.Name) {
			continue
		}

		content, err := io.ReadAll(tarReader)
		if err != nil {
			return nil, stacktrace.Wrap(err)
		}
		return content, nil
	}

	return nil, stacktrace.New("Error during tar extraction: no file with the name 'duckdb' found in the tar file.")
}

func extractFromZip(assetBytes []byte) ([]byte, stacktrace.Error) {
	zipReader, err := zip.NewReader(bytes.NewReader(assetBytes), int64(len(assetBytes)))
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}
	for _, file := range zipReader.File {
		fileName := file.Name

		// Normal release -> just one file with the name "duckdb"
		if isDuckdbBinary(fileName) {
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
