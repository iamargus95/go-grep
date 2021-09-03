package main

import (
	"flag"
	"fmt"
	"iamargus95/gogrep/gogrep"
	"iamargus95/gogrep/iofile"
	"log"
	"sync"
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

	var recursive bool
	flag.BoolVar(&recursive, "r", false, "Searches every file inside a directory.")

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
		go worker(caseSensitive, count, after, before, path, pattern, outputChan, &wg)
	}

	go func() {

		for output := range outputChan {
			for i := 0; i < len(output); i++ {
				fmt.Println(output[i])
			}
		}
	}()

	wg.Wait()
}

func worker(caseSensitive bool, count bool, after int, before int, path, pattern string,
	outputChan chan []string, wg *sync.WaitGroup) {

	var result []string
	defer wg.Done()

	fileContents, _ := iofile.ReadFile(path)

	if caseSensitive {
		result = gogrep.GrepCaseInsensitive(fileContents, pattern)

	} else if count {
		result = gogrep.GrepCount(fileContents, pattern)

	} else if after > 0 {
		result = gogrep.GrepAfter(after, fileContents, pattern)

	} else if before > 0 {
		result = gogrep.GrepBefore(before, fileContents, pattern)

	} else {
		result = gogrep.Grep(fileContents, pattern)
	}

	outputChan <- result
}
