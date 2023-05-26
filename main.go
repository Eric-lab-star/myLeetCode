package main
//pong
import "fmt"

/*
palindrome number
*/
func main() {
	input := 121
	output := isPalindrome(input)
	fmt.Println(output)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	digits := parser(x)
	for i, j := 0, len(digits)-1; i <= j; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}
	return true
}

func parser(x int) map[int]int {
	digits := make(map[int]int)
	for i := 0; ; i++ {
		digits[i] = (x % power(i+1)) / power(i)
		x = x - (x % power(i+1))
		if x < power(i) {
			break
		}
	}
	return digits
}
func power(x int) int {
	var output = 1
	for i := 0; i < x; i++ {
		output = output * (10)
	}
	return output
}
