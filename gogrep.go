package gogrep

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

func SearchString(path, pattern string) string {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	r, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			return (scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return (scanner.Text())

}

func PatternString() string {
	return (os.Args[2])
}
