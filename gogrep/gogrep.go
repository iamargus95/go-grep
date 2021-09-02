package gogrep

import (
	"strconv"
	"strings"
)

func GrepCaseInsensitive(fileContents []string, pattern string) []string {
	var outputString []string
	for index, line := range fileContents {
		lineLower := strings.ToLower(line)
		patternLower := strings.ToLower(pattern)
		if strings.Contains(lineLower, patternLower) {
			outputString = append(outputString, (fileContents[index]))
			outputString = append(outputString, "--")
		}
	}
	return outputString
}

func GrepCount(fileContents []string, pattern string) []string {
	var outputString []string
	var countOutput int
	var outputSlice []string

	for _, line := range fileContents {
		if strings.Contains(line, pattern) {
			countOutput += strings.Count(line, pattern)
		}
		outputStr := strconv.Itoa(countOutput)
		outputString = append(outputString, outputStr)
	}

	outputSlice = append(outputSlice, outputString[(len(outputString)-1)])
	return outputSlice
}

func Grep(fileContents []string, pattern string) []string {
	var outputString []string
	for index, line := range fileContents {
		if strings.Contains(line, pattern) {
			outputString = append(outputString, (fileContents[index]))
			outputString = append(outputString, "--")

		}
	}
	return outputString
}

func GrepAfter(after int, fileContents []string, pattern string) []string {
	var outputString []string

	for index, line := range fileContents {

		if strings.Contains(line, pattern) {
			length := len(fileContents)
			if index+after+1 > length {
				after = length - index - 1
			}

			for i := 0; i <= after; i++ {
				outputString = append(outputString, (fileContents[index+i]))
			}
			outputString = append(outputString, "--")
		}
	}
	return outputString
}

func GrepBefore(before int, fileContents []string, pattern string) []string {
	var outputString []string
	for index, line := range fileContents {

		if strings.Contains(line, pattern) {

			if index-before <= 0 {
				before = 0
			}

			for i := before; i <= index; i++ {
				outputString = append(outputString, (fileContents[i]))
			}
			outputString = append(outputString, "--")
		}
	}
	return outputString
}
