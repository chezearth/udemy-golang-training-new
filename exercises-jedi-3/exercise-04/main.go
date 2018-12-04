package main

import "fmt"

func main() {
	y := 1967
	for {
		if y > 2018 {
			break
		}
		fmt.Printf("%v ", y)
		y++
	}
	fmt.Println("")
}
