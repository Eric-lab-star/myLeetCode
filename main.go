package main

import (
	"fmt"
)

func main() {
	input := []int{11, 14, 2, 7}
	output := twoSum(input, 9)
	fmt.Println(output)
}

func twoSum(nums []int, target int) []int {
	var output []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				output := []int{i, j}
				return output
			}
		}
	}
	return output
}
