package main

import (
	"flag"
	"fmt"
	"iamargus95/gogrep/gogrep"
	"iamargus95/gogrep/iofile"
)

func main() {

	var caseSensitive bool
	flag.BoolVar(&caseSensitive, "i", false, "Do a Case-Insensitive Search.")

	var count bool
	flag.BoolVar(&count, "c", false, "Number of matches in a string.")

	var after int
	flag.IntVar(&after, "A", 0, "Shows number of lines after the Match.")

	var before int
	flag.IntVar(&before, "B", 0, "Shows number of lines before the Match.")

	flag.Parse()

	var fileContents []string
	pattern := flag.Arg(0)
	listOfFiles := iofile.ListFilesInDir(flag.Arg(1))
	for i := 0; i < len(listOfFiles); i++ {
		fileContents, _ = iofile.ReadFile(listOfFiles[i])
		if caseSensitive {
			output := gogrep.GrepCaseInsensitive(fileContents, pattern)

			for i := 0; i < len(output); i++ {
				fmt.Println(output[i])
			}

		} else if count {
			output := gogrep.GrepCount(fileContents, pattern)
			fmt.Println(output)

		} else if after > 0 {
			output := gogrep.GrepAfter(after, fileContents, pattern)

			for i := 0; i < len(output); i++ {
				fmt.Println(output[i])
			}

		} else if before > 0 {
			output := gogrep.GrepBefore(before, fileContents, pattern)

			for i := 0; i < len(output); i++ {
				fmt.Println(output[i])
			}

		} else {
			output := gogrep.Grep(fileContents, pattern)

			for i := 0; i < len(output); i++ {
				fmt.Println(output[i])
			}
		}
	}
}
