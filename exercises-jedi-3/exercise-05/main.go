package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("%v divided by 4, remainder is %v. ", i, i%4)
	}
	fmt.Println("")
}
