package main

import "fmt"

// main func is a multiplication algorithm of quadratic complexity that takes n*n time to complete i.e. O(n²) time complexity
func main() {
    var i, k int
    for k = 1; k <= 10; k++ {
        fmt.Println("Multiplying with: ", k)
        for i = 1; i <= 10; i ++ {
            var mult = k * i
            fmt.Println(mult)
        }
    }
}