package main

import "fmt"

func main() {
	ret := isPalindrome(14141)
	fmt.Println(ret)
}

func isPalindrome(num int) bool {
	if num < 0 {
		return false
	}

	divider := 1
	for {
		if num/(divider*10) == 0 {
			break
		}

		divider *= 10
	}

	for num > 0 {
		fmt.Println(divider)
		//앞자리수와 뒷자리수 비교
		if num/divider != num%10 {
			return false
		}
		//앞자리수와 뒷자리수 자르기
		//14141 -> 414
		num = num % divider
		num /= 10
		// 2자리수가 줄었음으로 100으로 나누기
		divider /= 100
	}

	return true
}
