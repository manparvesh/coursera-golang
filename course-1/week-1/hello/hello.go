package main

import "fmt"

// func main() {
// 	fmt.Println("Hello, world!")
// }
func main() {
  var x int
  var y *int
  z := 3
  y = &z
  x = &y
}
