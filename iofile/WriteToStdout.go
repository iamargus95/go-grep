package iofile

import "fmt"

func WriteToStdout(result chan []string) {

	length := len(result)

	for i := 0; i < length; i++ {
		output, open := <-result

		if !open {
			break
		}

		fmt.Println(output[i])
	}

}
