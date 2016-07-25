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

// fallback: this is an int32. large n causes overflow.
func fib(n int, receiver chan int) {
	previous := -1
	result := 1
	for i := 0; i < n; i++ {
		sum := result + previous
		previous = result
		result = sum
	}

	receiver <- result
}