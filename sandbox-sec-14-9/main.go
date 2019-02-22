package main

import "fmt"

func main() {
	fmt.Println(odd(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...))
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func odd(f func(xi ...int) int, yi ...int) int {
	var zi []int
	for _, v := range yi {
		if v%2 != 0 {
			zi = append(zi, v)
		}
	}
	return f(zi...)
}
