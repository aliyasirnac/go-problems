package main

import "fmt"

func f(n int) int {
	if n < 2 {
		return n
	}
	return f(n-2) + f(n-1)
}

func main() {
	index := 6
	fmt.Println(f(index))
}
