package main

import (
	"fmt"
)

func main() {
	input := []int{2}
	fmt.Printf("input %v\n", input)
	output := plusOne(input)
	fmt.Printf("output %v\n", output)
}

func plusOne(digits []int) []int {
	length := len(digits)
	for {
		if digits[length-1]+1 == 10 {
			digits[length-1] = 0
			if length-1 == 0 {
				output := []int{1}
				return append(output, digits...)
			}
			length--
		} else {
			digits[length-1]++
			return digits
		}

	}
}
