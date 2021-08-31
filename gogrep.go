package gogrep

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

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

	return txtlines, err
}

func Grep(fileContents []string, pattern string) []string {

	var outputString []string

	var caseSensitive bool
	flag.BoolVar(&caseSensitive, "i", false, "Do a Case-Insensitive Search.")

	var count bool
	flag.BoolVar(&count, "c", false, "Number of matches in a string.")

	flag.Parse()

	if caseSensitive {
		for index, line := range fileContents {
			lineLower := strings.ToLower(line)
			patternLower := strings.ToLower(pattern)
			if !strings.Contains(lineLower, patternLower) {

			} else {
				outputString = append(outputString, (fileContents[index]))
			}
		}
		return outputString

	} else if count {
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

	} else {

		for index, line := range fileContents {
			if !strings.Contains(line, pattern) {

			} else {
				outputString = append(outputString, (fileContents[index]))
			}
		}
		return outputString
	}
}
