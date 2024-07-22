package main

import (
	"fmt"
	"time"
)

// func main() {
// 	strs := make(chan string, 1000)
// 	t := time.Now()
// 	go task2(strs, "lol")
// 	for i := 0; i < cap(strs); i++ {
// 		fmt.Print(<-strs)
// 	}
// 	close(strs)
// 	fmt.Print(time.Since(t))
// }

// func task2(strs chan<- string, s string) {
// 	for i := 0; i < cap(strs); i++ {
// 		strs <- s + " "
// 	}
// }

func main() {
	strs := make(chan string)
	t := time.Now()
	go task2(strs, "lol")
	for i := 0; i < 1000; i++ {
		fmt.Print(<-strs)
	}
	fmt.Print(time.Since(t))
}
func task2(strs chan<- string, s string) {
	for i := 0; i < 1000; i++ {
		strs <- s + " "
	}
}
