package main

import (
	"fmt"
	"strings"
)

func main() {
	input2 := "a"
	out := lengthOfLastWord(input2)
	fmt.Println(out)
}

func lengthOfLastWord(s string) int {
	var out int
	var space = rune(32)
	var words []string
	var prevIndex int
	s = strings.TrimSpace(s)
	if rune(s[len(s)-1]) != space {
		s = s + string(space)
	}
	for index, runeValue := range s {
		if runeValue == space {
			words = append(words, s[prevIndex:index])

			prevIndex = index + 1
		}
	}
	out = len(words[len(words)-1])
	fmt.Println(words[len(words)-1])
	fmt.Println(words)
	return out
}

// "my name is kim" --> 3
