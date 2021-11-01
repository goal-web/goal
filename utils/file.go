package utils

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func CopyFile(from, to string, bufferSize int64) error {
	sourceFileStat, statRrr := os.Stat(from)
	if statRrr != nil {
		return statRrr
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file.", from)
	}

	source, openErr := os.Open(from)
	if openErr != nil {
		return openErr
	}
	defer source.Close()

	_, statRrr = os.Stat(to)
	if statRrr == nil {
		return fmt.Errorf("File %s already exists.", to)
	}

	destination, createErr := os.Create(to)
	if createErr != nil {
		return createErr
	}
	defer destination.Close()

	buf := make([]byte, bufferSize)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err = destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}

func AllFiles(path string) (results []fs.FileInfo) {
	_ = filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err == nil {
			results = append(results, info)
		}
		return nil
	})
	return results
}

func AllDirectories(path string) (results []string) {
	_ = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err == nil {
			results = append(results, d.Name())
		}
		return nil
	})
	return results
}
