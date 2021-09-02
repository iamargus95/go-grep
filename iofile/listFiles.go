package iofile

import (
	"log"
	"os"
	"path/filepath"
)

func ListFilesInDir(rootFilePath string) []string {
	inputFiles := []string{}
	err := filepath.Walk(rootFilePath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				inputFiles = append(inputFiles, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return inputFiles
}
