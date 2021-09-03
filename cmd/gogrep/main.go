package main

import (
	"flag"
	"iamargus95/gogrep/gogrep"
	"iamargus95/gogrep/io"
	"log"
	"sync"
)

func main() {

	var after int
	flag.IntVar(&after, "A", 0, "Shows number of lines after the Match.")

	var before int
	flag.IntVar(&before, "B", 0, "Shows number of lines before the Match.")

	var caseSensitive bool
	flag.BoolVar(&caseSensitive, "i", false, "Do a Case-Insensitive Search.")

	var count bool
	flag.BoolVar(&count, "c", false, "Number of matches in a string.")

	flag.Parse()

	pattern := flag.Arg(0)
	rootPath := flag.Arg(1)

	paths, err := io.ListFiles(rootPath)
	if err != nil {
		log.Fatal(err)
	}

	outputChan := make(chan []string)
	var wg sync.WaitGroup

	if caseSensitive {

		for _, path := range paths {
			wg.Add(1)
			go workerCase(caseSensitive, path, pattern, outputChan, &wg)
		}

	} else if after > 0 {

		for _, path := range paths {
			wg.Add(1)
			go workerAfter(after, path, pattern, outputChan, &wg)
		}

	} else if before > 0 {

		for _, path := range paths {
			wg.Add(1)
			go workerBefore(before, path, pattern, outputChan, &wg)
		}

	} else if count {

		for _, path := range paths {
			wg.Add(1)
			go workerCount(count, path, pattern, outputChan, &wg)
		}

	} else {

		for _, path := range paths {
			wg.Add(1)
			go worker(path, pattern, outputChan, &wg)
		}
	}

	go io.WriteToStdout(outputChan)
	wg.Wait()

}

func workerAfter(after int, path string, pattern string, outputChan chan []string, wg *sync.WaitGroup) {
	var result []string
	defer wg.Done()
	fileContents, _ := io.ReadFile(path)
	result = gogrep.GrepAfter(after, fileContents, pattern)
	outputChan <- result
}

func workerBefore(before int, path string, pattern string, outputChan chan []string, wg *sync.WaitGroup) {
	var result []string
	defer wg.Done()
	fileContents, _ := io.ReadFile(path)
	result = gogrep.GrepBefore(before, fileContents, pattern)
	outputChan <- result
}

func workerCase(caseSensitive bool, path string, pattern string, outputChan chan []string, wg *sync.WaitGroup) {
	var result []string
	defer wg.Done()
	fileContents, _ := io.ReadFile(path)
	result = gogrep.GrepCaseInsensitive(fileContents, pattern)
	outputChan <- result
}

func workerCount(count bool, path string, pattern string, outputChan chan []string, wg *sync.WaitGroup) {
	var result []string
	defer wg.Done()
	fileContents, _ := io.ReadFile(path)
	result = gogrep.GrepCount(fileContents, pattern)
	outputChan <- result
}

func worker(path string, pattern string, outputChan chan []string, wg *sync.WaitGroup) {
	var result []string
	defer wg.Done()
	fileContents, _ := io.ReadFile(path)
	result = gogrep.Grep(fileContents, pattern)
	outputChan <- result
}
