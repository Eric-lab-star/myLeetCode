package main

import (
	"fmt"
	"math"
)

func main() {
	digits := []int{1, 2, 4}
	int := marshal(digits)
	fmt.Println(int)
	upperInt := unmarshal(int + 1)
	fmt.Println(upperInt)
}

func plusOne(digits []int) []int {
	return unmarshal(marshal(digits) + 1)

}

func marshal(digits []int) int {
	var length = len(digits)
	var inputInt int
	for index, value := range digits {
		stat := math.Pow(10, float64(length-index-1))
		inputInt += value * int(stat)
	}
	return inputInt
}

func unmarshal(number int) []int {
	var out []int
	stringN := fmt.Sprintln(number)
	var length = len(stringN) - 1
	stat := int(math.Pow10(length - 1))
	for length != 0 {
		if number < 10 {
			out = append(out, number)
		} else {
			remain := number % stat
			out = append(out, number/stat)
			number = remain
		}
		length--
	}
	fmt.Println(out)
	return out

}
