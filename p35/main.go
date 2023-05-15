package main

import "fmt"

func main() {
	target := 5
	nums := []int{1, 3, 5, 6}
	ans := searchInsert(nums, target)
	fmt.Println(ans)
}
func searchInsert(nums []int, target int) int {
	for i, d := range nums {
		if d == target {
			return i
		} else if d > target {
			return i
		}
	}
	return len(nums)
}
