package main

import "fmt"

func main() {
	input := []string{"", "a"}
	ret := longestCommonPrefix(input)
	fmt.Println(ret)

}

func longestCommonPrefix(strs []string) string {
	var prefix string
	// only one item in the array
	if len(strs) == 1 {
		return strs[0]
	}
	for i := 0; i < len(strs); i++ {
		if strs[i] == "" {
			return ""
		}
		if prefix == "" {
			for j, char := range strs[i+1] {
				// cannot compare char
				if j == len(strs[i]) {
					break
				}
				if char != rune(strs[i][j]) {
					//no prefix
					if j == 0 {
						return ""
					}
					// same char until index j
					prefix = strs[i][:j]
					break
				}
				// all charters are same
				if len(strs[i]) < len(strs[i+1]) {
					prefix = strs[i]
					break
				}

				prefix = strs[i+1]

			}
		}

		if prefix != "" {
			if strs[i] == "" {
				return ""
			}
			for j, char := range prefix {
				// cannot compare char

				if j >= len(strs[i]) {
					//no prefix
					if j == 0 {
						return ""
					}
					break
				}
				if char != rune(strs[i][j]) {
					//no prefix
					if j == 0 {
						return ""

					}
					// same char until index j
					prefix = strs[i][:j]
					break
				}
				// all charters are same
				if len(strs[i]) < len(prefix) {
					prefix = strs[i]
					break
				}

			}
		}

	}
	return prefix

}

/*
어려운 점
1. 배열을 돌려는 것과 동시에 문자열의 캐릭터를 톨려야함
총 돌려야 하는 것들이 3개이다.

2. prefix가 없을 때 어떻게 할 것인가?

*/
