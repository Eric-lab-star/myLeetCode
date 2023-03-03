package main

import (
	"fmt"
)

func main() {
	INPUT := "XIV"
	ret := romanToInt(INPUT)
	fmt.Println(ret)

}

func romanToInt(s string) int {
	romanNumberDict := map[rune]uint16{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var ret uint16
	var pre uint16

	for _, char := range s {
		cur := romanNumberDict[char]
		if cur > pre {
			ret -= pre * 2
		}
		ret += cur
		pre = cur
	}

	return int(ret)
}
