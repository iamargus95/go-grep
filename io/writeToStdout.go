package io

import "fmt"

func WriteToStdout(output chan []string) {

	for output != nil {
		for print := range output {
			for i := 0; i < len(print); i++ {
				fmt.Println(print[i])
			}
		}
	}
}
