package main

import (
	"flag"
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

	result := make(chan []string)
	linesInFile := make(chan []string)

	pattern := flag.Arg(0)

	filesToBeRead := iofile.ListFilesInDir(flag.Arg(1))

	for i := 0; i < len(filesToBeRead); i++ {
		go iofile.ReadFile(filesToBeRead[i], linesInFile)

		if caseSensitive {
			go gogrep.GrepCaseInsensitive(<-linesInFile, pattern, result)

		} else if count {
			go gogrep.GrepCount(<-linesInFile, pattern, result)

		} else if after > 0 {
			go gogrep.GrepAfter(after, <-linesInFile, pattern, result)

		} else if before > 0 {
			go gogrep.GrepBefore(before, <-linesInFile, pattern, result)

		} else {
			go gogrep.Grep(<-linesInFile, pattern, result)

		}
	}

	iofile.WriteToStdout(result)
}
