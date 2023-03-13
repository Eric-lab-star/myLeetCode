package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []string{"flow", "flower", "flight"}
	ret := longestCommonPrefix(input)
	fmt.Println(ret)
}

func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	maxPrefix := 0
	first := strs[0]
	last := strs[len(strs)-1]
	for i := range first {
		if first[i] != last[i] {
			break
		}
		maxPrefix++
	}
	return first[:maxPrefix]
}

/*
어려운 점
1. 배열을 돌려는 것과 동시에 문자열의 캐릭터를 톨려야함
총 돌려야 하는 것들이 3개이다.

2. prefix가 없을 때 어떻게 할 것인가?

*/
