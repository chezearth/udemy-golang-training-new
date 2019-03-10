package main

import "fmt"

func main() {
	t := bar(sum, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}...)
	fmt.Println(t)

	linkUp(concat, []string{"fred", " ", "john", " ", "luca"}...)
}

func bar(f func(x ...int) int, y ...int) int {
	var s []int
	for _, v := range y {
		if v%2 == 0 {
			s = append(s, v)
		}
	}
	total := f(s...)

	return total
}

func sum(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func concat(x ...string) string {
	var n string
	for _, v := range x {
		n = n + v
	}
	return n
}

func linkUp(f func(ss ...string) string, s ...string) {
	var si []string
	for _, v := range s {
		si = append(si, v)
	}
	result := f(si...)
	fmt.Println(result)
}
