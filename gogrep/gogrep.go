package gogrep

import (
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
	var outputSlice []string

	for _, line := range fileContents {
		if !strings.Contains(line, pattern) {

		} else {
			countOutput += strings.Count(line, pattern)
		}
		outputStr := strconv.Itoa(countOutput)
		outputString = append(outputString, outputStr)
	}

	outputSlice = append(outputSlice, outputString[(len(outputString)-1)])
	return outputSlice
}

func Grep(fileContents []string, pattern string) []string {

	for index, line := range fileContents {

		if !strings.Contains(line, pattern) {

		} else {
			outputString = append(outputString, (fileContents[index]))
		}
	}
	return outputString
}
