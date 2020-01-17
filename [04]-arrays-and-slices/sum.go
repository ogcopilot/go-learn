package main

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
