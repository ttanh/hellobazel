package main

import "fmt"

func main() {
	fmt.Printf("Sum of %d and %d is %d \n", 10, 20, sum(10, 20))
}

func sum(a, b int) int {
	return a + b
}
