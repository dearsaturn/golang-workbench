package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	receiver := make(chan int, 1)
	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(err)
	}

	go fib(n, receiver)
	fmt.Println(<-receiver)
}

func fib(n int, receiver chan int) {
	previous, result := -1, 1

	for i := 0; i < n; i++ {
		previous, result = result, result+previous
	}

	receiver <- result
	close(receiver)
}
