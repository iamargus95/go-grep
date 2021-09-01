package gogrep

import (
	"fmt"
	"strconv"
	"strings"
)

var outputString []string

func GrepCaseInsensitive(fileContents []string, pattern string) []string {

	for index, line := range fileContents {
		lineLower := strings.ToLower(line)
		patternLower := strings.ToLower(pattern)
		if !strings.Contains(lineLower, patternLower) {

		} else {
			outputString = append(outputString, (fileContents[index]))
		}
	}
	return outputString
}

func GrepCount(fileContents []string, pattern string) []string {

	var countOutput int
	var length int
	var outputSlice []string

	for _, line := range fileContents {
		if !strings.Contains(line, pattern) {

		} else {
			countOutput += strings.Count(line, pattern)
		}
		outputStr := strconv.Itoa(countOutput)
		outputString = append(outputString, outputStr)
	}

	length = len(outputString) - 1
	outputSlice = append(outputSlice, outputString[length])
	return outputSlice
}

func Grep(fileContents []string, pattern string) []string {

	for index, line := range fileContents {
		if !strings.Contains(line, pattern) {

		} else {
			outputString = append(outputString, (fileContents[index]))
		}
	}
	fmt.Println(outputString)
	return outputString
}
