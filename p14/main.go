package main

import (
	"fmt"
	"strings"
)

func main() {
	input := []string{"flower", "flow", "flight"}
	ret := longestCommonPrefix(input)
	fmt.Println(ret)
}

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		fmt.Println(prefix)
		if !strings.HasPrefix(strs[i], prefix) {
			prefix = ""
		}
	}
	return prefix

}

// 1. flower를 flow와
// 2. prefix와 flow를 한글자 한글자 단위로 비교
// 	2-1 e 부터 글자가 다름으로, prefix 변경
// 3. 반복
