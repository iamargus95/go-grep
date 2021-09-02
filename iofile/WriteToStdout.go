package iofile

import "fmt"

func WriteToStdout(result chan []string) {
	output := <-result
	for i := 0; i < len(output); i++ {
		fmt.Println(output[i])
	}
}
