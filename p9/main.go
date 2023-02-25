package main

import "fmt"

func main() {
	input := 0

	ans := isPalindrome(input)

	fmt.Println(ans)

}
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	a := IntToSlice(x, []int{})
	s := 0
	e := len(a) - 1

	for s <= (len(a)-1)/2 {
		if a[s] == a[e] {
			s++
			e--
			continue
		}
		return false
	}

	return true
}

func IntToSlice(n int, sequence []int) []int {
	if n != 0 {
		i := n % 10
		sequence = append([]int{i}, sequence...)
		return IntToSlice(n/10, sequence)
	}
	return sequence
}
