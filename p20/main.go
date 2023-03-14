package main

import "fmt"

func main() {
	ret := isValid("([]")
	fmt.Println(ret)

}

func isValid(s string) bool {
	var OpenParenthese = map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	var keep = []rune{}
	//iterate over input s
	for i := 0; i < len(s); i++ {
		char := rune(s[i])
		closedParenthese, ok := OpenParenthese[char]
		// if char is one of open parenthese
		if ok {
			// keep in closed parenthese slice
			keep = append(keep, closedParenthese)
		} else if len(keep) > 0 {
			// if char is not open parenthese, check last elem of "keep" slice
			// return false if does not match
			if keep[len(keep)-1] != char {
				return false
			}
			// cut last element
			keep = keep[:len(keep)-1]
		} else {
			return false
		}
	}
	// if lengthe of keep is greater than 0 return false
	//
	if len(keep) > 0 {
		return false
	}
	return true
}

/*
* "[[" 일때 오류 생김,
*
Open brackets must be closed by the same type of brackets. -> 쌍
Open brackets must be closed in the correct order. -> 순서
Every close bracket has a corresponding open bracket of the same type. -> 쌍
*/
