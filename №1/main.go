package main

import (
	"fmt"
)

func main() {
	numbers := make(chan int)
	go task(numbers, 5)
	fmt.Print(<-numbers)
}

func task(numbers chan int, n int) {
	numbers <- n + 1
}
