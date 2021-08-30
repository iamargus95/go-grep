package main

import (
	"flag"
	"fmt"
	"os"

	gogrep "github.com/iamargus95/go-grep/gogrep"
)

func main() {

	var after int
	flag.IntVar(&after, "A", 2, "number of lines after the match.")

	var before int
	flag.IntVar(&before, "B", 2, "number of lines before the match.")

	var around int
	flag.IntVar(&around, "C", 2, "number of lines before and after the match.")

	var recursive bool
	flag.BoolVar(&recursive, "r", true, "List files in a directory recursively.")

	var insensitive bool
	flag.BoolVar(&insensitive, "i", false, "Case-Insensitive Search.")

	flag.Parse()

	pattern := gogrep.PatternString()
	path := os.Args[2]
	fmt.Printf(gogrep.SearchString(path, pattern) + "\n")
}
