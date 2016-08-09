package fib

func Fib(n int) int {
	return <-fib(n)
}

func fib(n int) (result_chan chan int) {
	result_chan = make(chan int, 1)

	go func() {
		previous, result := -1, 1
		for i := 0; i < n; i++ {
			previous, result = result, result+previous
		}

		result_chan <- result
		close(result_chan)
	}()

	return
}
