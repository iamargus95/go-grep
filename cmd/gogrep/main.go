package main

import (
	"flag"
	"iamargus95/gogrep/gogrep"
	"iamargus95/gogrep/iofile"
	"log"
	"sync"
)

func main() {

	flag.Parse()

	pattern := flag.Arg(0)
	rootPath := flag.Arg(1)

	paths, err := iofile.ListFiles(rootPath)
	if err != nil {
		log.Fatal(err)
	}

	outputChan := make(chan []string)
	var wg sync.WaitGroup
	for _, path := range paths {
		wg.Add(1)
		go worker(path, pattern, outputChan, &wg)

	}

	go func() {

		iofile.WriteToStdout(outputChan)

	}()

	wg.Wait()
}

func worker(path, pattern string, outputChan chan<- []string, wg *sync.WaitGroup) {

	defer wg.Done()

	var caseSensitive bool
	flag.BoolVar(&caseSensitive, "i", false, "Do a Case-Insensitive Search.")

	var count bool
	flag.BoolVar(&count, "c", false, "Number of matches in a string.")

	var after int
	flag.IntVar(&after, "A", 0, "Shows number of lines after the Match.")

	var before int
	flag.IntVar(&before, "B", 0, "Shows number of lines before the Match.")

	fileContents, _ := iofile.ReadFile(path)

	if caseSensitive {
		gogrep.GrepCaseInsensitive(fileContents, pattern)

	} else if count {
		gogrep.GrepCount(fileContents, pattern)

	} else if after > 0 {
		gogrep.GrepAfter(after, fileContents, pattern)

	} else if before > 0 {
		gogrep.GrepBefore(before, fileContents, pattern)

	} else {
		gogrep.Grep(fileContents, pattern)
	}
}
