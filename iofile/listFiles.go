package iofile

import (
	"bufio"
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

func ReadFile(filepath string, linesInFile chan []string) {

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

	linesInFile <- txtlines
	close(linesInFile)
}
