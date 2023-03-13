package main

import "fmt"

func main() {
	ret := isValid("([)")
	fmt.Println(ret)

}

func isValid(s string) bool {
	var parenthese = map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	var keep = []rune{}
	if len(s)%2 == 1 {
		return false
	}
	for i := 0; i < len(s); i++ {
		fmt.Println(i)
		fmt.Println(string(keep))
		v, ok := parenthese[rune(s[i])]

		if ok {
			keep = append(keep, v)
			continue
		}
		if !ok && keep[len(keep)-1] == rune(s[i]) {
			keep = keep[:len(keep)-1]
			continue
		}
		if !ok && keep[len(keep)-1] != rune(s[i]) {
			return false

		}
	}
	return true
}
