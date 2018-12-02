package main

import "fmt"

func main()  {
  const (
    x1 = 2018 - iota
    x2 = 2018 - iota
    x3 = 2018 - iota
    x4 = 2018 - iota
  )
  fmt.Println(x1)
  fmt.Println(x2)
  fmt.Println(x3)
  fmt.Println(x4)
}
