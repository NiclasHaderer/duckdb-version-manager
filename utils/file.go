package utils

import (
	"duckdb-version-manager/stacktrace"
	"io"
	"os"
	"path/filepath"
)

func RemoveFileOrDie(file string) {
	if _, err := os.Stat(file); err == nil {
		err := os.RemoveAll(file)
		if err != nil {
			ExitWithError(stacktrace.Wrap(err))
		}
	}
}

// WriteFileAtomic replaces path via rename instead of truncating it, so readers never see a partial file
// and an executable that is currently running (ETXTBSY on Linux) can still be replaced.
func WriteFileAtomic(path string, data []byte, perm os.FileMode) stacktrace.Error {
	tmp, err := os.CreateTemp(filepath.Dir(path), filepath.Base(path)+".*.tmp")
	if err != nil {
		return stacktrace.Wrap(err)
	}
	defer os.Remove(tmp.Name())

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return stacktrace.Wrap(err)
	}
	if err := tmp.Chmod(perm); err != nil {
		tmp.Close()
		return stacktrace.Wrap(err)
	}
	if err := tmp.Close(); err != nil {
		return stacktrace.Wrap(err)
	}
	return stacktrace.Wrap(os.Rename(tmp.Name(), path))
}

func CopyFile(src, dst string) stacktrace.Error {
	srcFile, err := os.Open(src)
	if err != nil {
		return stacktrace.Wrap(err)
	}

	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return stacktrace.Wrap(err)
	}

	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return stacktrace.Wrap(err)
	}

	return nil
}
