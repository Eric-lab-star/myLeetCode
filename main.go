package main

import (
	"fmt"
)

func main() {
	a := "111"
	b := "1"
	output := addBinary(a, b)
	fmt.Println(output)
}

func addBinary(a string, b string) string {
	sA := toSlice(a)
	sB := toSlice(b)
	extraOne := false

	for i := len(sA) - 1; i >= 0; i-- {
		if extraOne && sA[i] == 49 {
			sA[i] = 48
		}
		if sA[i] == 49 && sB[0] == 49 {
			sA[i] = 48
			extraOne = true
		} else if sA[i] == 49 && sB[0] == 48 || sA[i] == 48 && sB[0] == 49 {
			sA[i] = 49
		} else {
			sA[i] = 48
		}

	}
	return string(sA)

}

func toSlice(i string) (slice []byte) {
	for _, v := range i {
		slice = append(slice, byte(v))
	}
	return

}

/*


 */
