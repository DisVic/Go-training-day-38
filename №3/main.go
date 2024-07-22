package main

import (
	"fmt"
	"time"
)

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)

	go removeDuplicates(inputStream, outputStream)
	go printer(outputStream)

	for _, i := range "112334456" {
		inputStream <- string(i)
	}

	var input string
	fmt.Scanln(&input)
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func removeDuplicates(inputStream, outputStream chan string) {
	defer close(outputStream)
	value1 := <-inputStream
	outputStream <- value1
	func() {
		for x := range inputStream {
			value2 := x
			if value1 != value2 {
				outputStream <- value2
			}
			value1 = x
		}
	}()
}
