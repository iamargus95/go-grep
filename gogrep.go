package gogrep

import (
	"bufio"
	"flag"
	"log"
	"os"
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

	caseSensitive := flag.Bool("i", true, "Do a Case-Insensitive Search.")
	flag.Parse()

	if *caseSensitive {
		for index, line := range fileContents {
			lineLower := strings.ToLower(line)
			patternLower := strings.ToLower(pattern)
			if !strings.Contains(lineLower, patternLower) {

			} else {
				outputString = append(outputString, (fileContents[index]))
			}
		}
		return outputString

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
