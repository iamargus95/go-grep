package io

import (
	"bufio"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func ListFiles(dir string) ([]string, error) {
	filePaths := []string{}
	err := filepath.WalkDir(dir,
		func(path string, d fs.DirEntry, err error) error {
			if !d.IsDir() {
				filePaths = append(filePaths, path)
			}
			return nil
		})
	if err != nil {
		return nil, err
	}

	return filePaths, nil
}

func ReadFile(filepath string) ([]string, error) {

	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	return txtlines, scanner.Err()
}
