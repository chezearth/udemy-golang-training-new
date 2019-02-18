package main

import "fmt"

type config struct {
	url     string
	wait    int
	retries int
}

func main() {
	c1 := config{"https://example.com/api/v1/users", 1000, 3}
	c2 := config{"https://example.com/api/v1/resources", 2000, 5}
	fmt.Println(c1.connect())
	fmt.Println(c2.connect())
}

func (c config) connect() string {
	s := fmt.Sprint("URL is ", c.url, ", wait interval is (ms) ", c.wait)
	return s
}
