package main

import "fmt"

func sumAll(number ...int) int {
	total := 0
	for _, value := range number {
		total += value
	}
	return total
}

func main() {
	total := sumAll(10, 10)
	fmt.Println(total)

	// slice parameter
	numbers := []int{10, 14, 1, 80}
	total = sumAll(numbers...)
	fmt.Println(total)
}
