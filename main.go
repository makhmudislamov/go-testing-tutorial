package main

import "fmt"

// Calculate returns input + 2
func Calculate (x int) (result int) {
	result = x + 2
	return result
}

func main() {
	fmt.Println("Running Test Tutorial Func")
	result := Calculate(2)
	fmt.Println(result)
}