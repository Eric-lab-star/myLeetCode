package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var Root = TreeNode{
	Val:   27,
	Left:  nil,
	Right: &ThirtyFour,
}

var Fifity = TreeNode{
	Val:   50,
	Left:  nil,
	Right: &FourtyFour,
}
var FourtyFour = TreeNode{
	Val:   44,
	Left:  nil,
	Right: nil,
}

var FiftyEight = TreeNode{
	Val:   58,
	Left:  &Fifity,
	Right: nil,
}

var ThirtyFour = TreeNode{
	Val:   34,
	Left:  nil,
	Right: &FiftyEight,
}

func main() {
	ret := minDiffInBST(&Root)
	fmt.Println(ret)
}

func minDiffInBST(root *TreeNode) int {

	minDiff := math.MaxInt32
	var preValue = -1
	inOrderTraversal(root, &preValue, &minDiff)

	return int(minDiff)

}

func inOrderTraversal(root *TreeNode, preValue *int, minDiff *int) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left, preValue, minDiff)

	if *preValue != -1 && root.Val-*preValue < *minDiff {
		*minDiff = root.Val - *preValue
	}

	*preValue = root.Val

	inOrderTraversal(root.Right, preValue, minDiff)

}

/*
mistakes
1. math.Abs(a,b float64)
2. Don't use Global variable on Leetcode
*/
