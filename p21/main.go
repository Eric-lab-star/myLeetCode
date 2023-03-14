package main

import (
	"fmt"
)

func main() {
	fmt.Println("heelo")

	ret := []int{}
	fmt.Println(ret)
}

/**
 * Definition for singly-linked list.
 You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
	}
	if list1 != nil {
		current.Next = list2
	} else {
		current.Next = list1
	}
	return result.Next

}
