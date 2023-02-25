package main

import "fmt"

// Time Limit Exceeded
func main() {
	ret := maxProfit(input)
	fmt.Println(ret)

}

var input = []int{7, 1, 5, 3, 6, 4}

func maxProfit(prices []int) int {
	var max int
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] > prices[i] {
				test := prices[j] - prices[i]
				if test > max {
					max = test
				}
			}
			continue

		}
	}
	return max
}
