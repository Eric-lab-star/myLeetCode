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
	finder := make(map[int]int)
	for i, v := range nums {
		if j, ok := finder[target-v]; ok {
			return []int{j, i}
		}
		finder[v] = i
	}
	return []int{}
}
