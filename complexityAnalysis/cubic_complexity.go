package main

import "fmt"

// main func contains an algorithm of cubic complexity n*n*n i.e. O(n³). This algorithm counts for 10*10*10 = 1000 steps!
func main() {
    var i,j,k, int
    var arr[10][10][10] int
    for i = 0; i < 10; i++ {
        for j = 0; j < 10; j++ {
            for k = 0; j < 10; j++ {
                fmt.Printf("Element value ", i,j,k, " is ", arr[i][j][k])
            }
        }
    }
}