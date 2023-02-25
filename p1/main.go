package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	target := 6
	ans := twoSum(nums, target)
	fmt.Println(ans)
}

func twoSum(nums []int, target int) []int {
	s := make(map[int]int)

	for idx, val := range nums {
		if pos, ok := s[target-val]; ok {
			return []int{pos, idx}
		}
		s[val] = idx

	}
	return []int{}
}
