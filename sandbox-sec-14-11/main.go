package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
}

func factorial(n int) int {
	result := 1
	for ; n > 0; n-- {
		result *= n
	}
	return result
}
