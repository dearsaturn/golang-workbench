package main

import "fmt"

func meat(name string) {
	fmt.Println(name)
}

func sandwich(fn func(content string)) func(content string) {
	return func(content string) {
		fmt.Println("--- Bread ---")
		fn(content)
		fmt.Println("--- Bread ---")
	}
}

func garnish(fn func(content string)) func(content string) {
	return func(content string) {
		fmt.Println("--- Olive ---")
		fn(content)
	}
}

func main() {
	garnish(sandwich(meat))("--- Salami ---")
}
