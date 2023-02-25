package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	target := 6
	ans := twoSum(nums, target)
	fmt.Println(ans)
}

func twoSum(nums []int, target int) []int {
	ret := []int{}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {

			if nums[i]+nums[j] == target && i != j {
				ret = []int{i, j}
				return ret
			}
		}
	}
	return ret
}
