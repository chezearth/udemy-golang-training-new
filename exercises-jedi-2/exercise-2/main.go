package main

import "fmt"

var a int = 2
var b = 10
var c = 10

func main() {
  expr0 := a == b
  expr1 := b == c
  expr2 := a <= b
  expr3 := b <= c
  expr4 := a >= b
  expr5 := b >= c
  expr6 := a != b
  expr7 := b != c
  expr8 := a < b
  expr9 := b < c
  expr10 := a > b
  expr11 := b > c

  fmt.Printf("%v %v %v %v %v %v %v %v %v %v %v %v\n", expr0, expr1, expr2, expr3, expr4, expr5, expr6, expr7, expr8, expr9, expr10, expr11)

}
