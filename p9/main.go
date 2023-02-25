package main

import "fmt"

func main() {
	input := -131

	ans := isPalindrome(input)

	fmt.Println(ans)

}
func isPalindrome(x int) bool {
	x_copy := x
	reverseX := 0

	for x_copy > 0 {
		reverseX = reverseX*10 + x_copy%10
		x_copy /= 10
	}

	return reverseX == x

}
