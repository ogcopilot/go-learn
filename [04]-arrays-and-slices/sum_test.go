package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func ExampleSum() {
	summed := Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(summed)
	// Output :45
}
