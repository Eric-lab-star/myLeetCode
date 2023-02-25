package main

import "fmt"

func main() {
	ret := isPalindrome(1234)
	fmt.Println(ret)
}

func isPalindrome(x int) (ans bool) {
	if x < 0 {
		return
	}
	y := 0
	originalNum := x
	for x > 0 {
		// y * 10은 y에다 0을 하나씩 붙여서 자리수를 증가시킴
		// x %10은 x의 1의 자리 숫자를 구해서 y에 붙여준다.
		y = y*10 + x%10
		// x/=10은 x에서 1의 자리 숫자를 제거 시킨다.
		x /= 10
	}
	// for 루프를 반복하면 1의 자리 숫자가 10의 자리 숫자로, 100의 자리 숫자로, 점점 진다. 결국 처음 인풋의 역순을 y에 할당해주게된다.
	fmt.Println(y)
	return originalNum == y
}
