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
		for results := range outputChan {
			fmt.Println(results)
		}
	}()

	wg.Wait()
}

func worker(path, pattern string, outputChan chan<- []string, wg *sync.WaitGroup) {

	defer wg.Done()

	var (
		caseSensitive bool
		count         bool
		after         int
		before        int
	)

	flag.BoolVar(&caseSensitive, "i", false, "Do a Case-Insensitive Search.")
	flag.BoolVar(&count, "c", false, "Number of matches in a string.")
	flag.IntVar(&after, "A", 0, "Shows number of lines after the Match.")
	flag.IntVar(&before, "B", 0, "Shows number of lines before the Match.")

	fileContents, _ := iofile.ReadFile(path)

	if caseSensitive {
		result := gogrep.GrepCaseInsensitive(fileContents, pattern)
		outputChan <- result

	} else if count {
		result := gogrep.GrepCount(fileContents, pattern)
		outputChan <- result

	} else if after > 0 {
		result := gogrep.GrepAfter(after, fileContents, pattern)
		outputChan <- result

	} else if before > 0 {
		result := gogrep.GrepBefore(before, fileContents, pattern)
		outputChan <- result

	} else {
		result := gogrep.Grep(fileContents, pattern)
		outputChan <- result
	}
}
