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

func GrepBefore(before int, fileContents []string, pattern string) []string { //----------> Does not handle multiple matches on separate strings which are
	var outputString []string               //----------------------------------------------closer than IntVar as expected. When adding "--"
	for index, line := range fileContents { //----------------------------------------------And reverses the order in which match was found.

		if strings.Contains(line, pattern) {
			outputString = append(outputString, "--")

			outputString = append(outputString, fileContents[index])

			if index-before <= 0 {
				before = index
			}

			for i := 1; i < before+1; i++ {
				outputString = append(outputString, (fileContents[index-i]))
			}
		}
	}
	for i, j := 0, len(outputString)-1; i <= j; i, j = i+1, j-1 {
		outputString[i], outputString[j] = outputString[j], outputString[i]
	}
	return outputString
}
