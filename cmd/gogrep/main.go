package main

import (
	"fmt"
	"os"

	gogrep "github.com/iamargus95/go-grep/gogrep"
)

func main() {

	fileContents, _ := gogrep.ReadFile(os.Args[3])
	pattern := os.Args[2]
	output := gogrep.Grep(fileContents, pattern)

	fmt.Println(output)
}
