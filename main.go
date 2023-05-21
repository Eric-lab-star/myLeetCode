package main

import "fmt"

func main() {
	input := 2147395599
	output := mySqrt(input)
	fmt.Println(output)
}

func mySqrt(x int) int {
	z := float64(1)
	floatX := float64(x)
	z = z - (z*z-floatX)/(2*z)
	for {
		newZ := z - (z*z-floatX)/(2*z)
		if z-newZ < 0.00001 {
			return int(z)
		}
		z = newZ
	}
}
