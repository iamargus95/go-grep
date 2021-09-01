package main

import (
	"flag"
	"fmt"

	gogrep "github.com/iamargus95/go-grep/gogrep"
	iofile "github.com/iamargus95/go-grep/iofile"
)

func main() {

	var caseSensitive bool
	flag.BoolVar(&caseSensitive, "i", false, "Do a Case-Insensitive Search.")

	var count bool
	flag.BoolVar(&count, "c", false, "Number of matches in a string.")

	flag.Parse()

	fileContents, _ := iofile.ReadFile(flag.Arg(1))
	pattern := flag.Arg(2)

	if caseSensitive {
		output := gogrep.GrepCaseInsensitive(fileContents, pattern)
		fmt.Println(output)
	} else if count {
		output := gogrep.GrepCount(fileContents, pattern)
		fmt.Println(output)
	} else {
		output := gogrep.Grep(fileContents, pattern)
		fmt.Println(output)
	}
}
