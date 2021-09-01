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

	flag.Parse()

	fileContents, _ := iofile.ReadFile(flag.Arg(1))
	pattern := flag.Arg(0)

	if caseSensitive {
		output := gogrep.GrepCaseInsensitive(fileContents, pattern)
		fmt.Println(output)
	} else if count {
		output := gogrep.GrepCount(fileContents, pattern)
		fmt.Println(output)
	} else if after != 0 {
		output := gogrep.GrepAfter(after, fileContents, pattern)
		fmt.Println(output)
	} else {
		output := gogrep.Grep(fileContents, pattern)
		fmt.Println(output)
	}
}
