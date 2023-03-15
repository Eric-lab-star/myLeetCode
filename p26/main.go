package main

import "fmt"

func main() {

	input := []int{1, 1, 2}
	k := removeDuplicates(input)
	fmt.Println(k)
	fmt.Println(input)

}

//  - - - -
// last = 2
// index = 1
// 1 2 2 2

// index = 1
// 1 2 2 3

// index = 3
// 1 2 2 3

func removeDuplicates(nums []int) int {
	last := nums[len(nums)-1]
	index := 1

	for index < len(nums) {
		if nums[index] == nums[index-1] {
			delete(nums, index)
			nums[len(nums)-1] = last + 1
			index--
		}
		if nums[index] == last+1 {

			break
		}
		index++
	}

	return index
}

func delete(nums []int, index int) {
	front := nums[:index]  // 1 2
	back := nums[index+1:] // 2
	for _, v := range back {
		front = append(front, v)
	}

}
