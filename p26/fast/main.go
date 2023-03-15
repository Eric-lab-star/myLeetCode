package main

import "fmt"

func main() {

	input := []int{1, 1, 2, 2, 3, 3}
	k := removeDuplicates(input)
	fmt.Println(k)
	fmt.Println(input)

}

func removeDuplicates(nums []int) int {
	i := 0
	j := 0

	for j != len(nums) {
		if nums[i] != nums[j] {
			nums[i+1], nums[j] = nums[j], nums[i+1]
			i++
		}
		j++
	}
	return i + 1
}
