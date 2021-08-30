package gogrep

import (
	"io/ioutil"
	"regexp"
)

func ReadFile(filepath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filepath)
	return data, err
}

func Grep(fileContents []byte, pattern string) []string {
	r := regexp.MustCompile(pattern)

	return (r.FindAllString(string(fileContents), -1))
}
