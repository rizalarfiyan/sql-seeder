package utils

import (
	"os"
	"sort"
)

func ReadDirectory(directory string) ([]os.FileInfo, error) {
	f, err := os.Open(directory)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	files, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}

	sort.Sort(NumericFilename(files))
	return files, nil
}
