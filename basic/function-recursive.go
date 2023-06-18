package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// recrusive function

func factiorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factiorialRecursive(value-1)
	}
}

func main() {
	loop1 := factorialLoop(4)
	loop := factiorialRecursive(4)
	fmt.Println(loop1)
	fmt.Println(loop)
	// fmt.Println(5 * 4 * 3 * 2 * 1)
}