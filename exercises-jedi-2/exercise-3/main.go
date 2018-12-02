package main

import (
  "fmt"
)

func main() {
  const (
    x int64 = 12768943202367
    y = 2345678
  )
  fmt.Printf("%T %d \t%T %d\n", x, x, y, y)
}
