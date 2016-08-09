package fib

import "testing"

func TestFib(t *testing.T) {
	want, got := 2, Fib(4)

	if want != got {
		t.Error("Expected 2, got", got)
	}
}
