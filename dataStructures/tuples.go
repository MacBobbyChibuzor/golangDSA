package main
// importing fmt package
import (
  "fmt"
)
//gets the power series of integer a and returns tuple of square of a
// and cube of a
func powerSeries(a int) (int,int) {
  return a*a, a*a*a
}

func main() {
	var square, cube int
	square, cube = powerSeries(3)
	fmt.Println(square, cube)
}