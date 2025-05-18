package myfib

import "fmt"

func Fib(n int) int {
	fib := []int{0, 1}

	for i := 2; i <= n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	fmt.Println("fib: ", fib)
	return fib[n]
}
