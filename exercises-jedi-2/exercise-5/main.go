package main

import "fmt"

func main()  {
  a := `This is a
  raw string
    literal,
      "indented".
This line starts at the margin.`
    fmt.Println(a)
}
