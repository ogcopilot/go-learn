package main

import "fmt"

// Sum will add all numbers pass in and return the total
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll will add all the numbers passed in batches and return one new batch of summed numbers
func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails will add all the numbers (but the first one) in batches and return one new batch with these partially summed numbers
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
