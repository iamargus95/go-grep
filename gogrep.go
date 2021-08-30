package gogrep

import (
	"io/ioutil"
	"strings"
)

func ReadFile(filepath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filepath)
	return data, err
}

func Grep(fileContents []byte, pattern string) []string {
	splitFile := strings.Split(string(fileContents), ".")
	var outputString []string
	for index, line := range splitFile {
		if !strings.Contains(line, pattern) {

		} else {
			outputString = append(outputString, (splitFile[index])+"\n")
		}
	}
	return outputString
}
