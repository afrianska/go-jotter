package main

import "fmt"

func sumAll(numbers ...int) int  {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	fmt.Println(sumAll(9, 4))
	
	slice := []int {10, 23, 55, 67, 45}
	fmt.Println(sumAll(slice...))

}