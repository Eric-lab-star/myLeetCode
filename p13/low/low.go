package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))

}

const roman = "IVXLCDM"

var nums = [7]uint16{1, 5, 10, 50, 100, 500, 1000}

func romanToInt(s string) int {
	var (
		sum      uint16
		previous uint16
	)

	for i := range s {
		v := uint16(0)
		for j := range roman {
			if s[i] == roman[j] {
				v = nums[j]
				break
			}
		}
		if v > previous {
			sum -= 2 * previous
		}
		sum += v
		previous = v
	}
	return int(sum)
}
