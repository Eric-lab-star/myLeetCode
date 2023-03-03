package main

import "fmt"

func romanToInt(s string) int {
	romanian := map[rune]int{
		'Z': 0,
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0
	for i, letter := range s {
		cur := romanian[letter]
		if (len(s) != i+1) && romanian[rune(s[i+1])] > cur {
			res -= cur
			continue
		}

		res += cur
	}

	return res
}

func main() {
	ret := romanToInt("III")
	fmt.Println(ret)
}
