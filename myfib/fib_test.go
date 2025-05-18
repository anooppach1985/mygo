package myfib

import (
	"testing"
)

func TestFib(t *testing.T) {
	if Fib(1) != 1 {
		t.Error("Expected 1, got ", Fib(1))
	}
	if Fib(2) != 1 {
		t.Error("Expected 1, got ", Fib(2))
	}
	if Fib(3) != 2 {
		t.Error("Expected 2, got ", Fib(3))
	}
	if Fib(4) != 3 {
		t.Error("Expected 3, got ", Fib(4))
	}
	if Fib(5) != 5 {
		t.Error("Expected 5, got ", Fib(5))
	}
	if Fib(6) != 8 {
		t.Error("Expected 8, got ", Fib(6))
	}
	if Fib(7) != 13 {
		t.Error("Expected 13, got ", Fib(7))
	}
	if Fib(8) != 21 {
		t.Error("Expected 21, got ", Fib(8))
	}
	if Fib(9) != 34 {
		t.Error("Expected 34, got ", Fib(9))
	}
	if Fib(10) != 55 {
		t.Error("Expected 55, got ", Fib(10))
	}
}
