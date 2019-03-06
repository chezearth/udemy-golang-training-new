package main

import "fmt"

func main() {
	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(foo(x...))
	fmt.Println(bar(x))
}

func foo(n ...int) int {
	result := 0
	for _, v := range n {
		result += v
	}
	return result
}

func bar(n []int) int {
	result := 0
	for _, v := range n {
		result += v
	}
	return result
}
