package main

import (
	"fmt"
)

var l1 = ListNode{1, &l1_1}
var l1_1 = ListNode{3, &l1_2}
var l1_2 = ListNode{5, nil}

var l2 = ListNode{2, &l2_1}
var l2_1 = ListNode{4, &l2_2}
var l2_2 = ListNode{6, nil}

func main() {
	merge := ListNode{}
	l1v := values(l1)
	l2v := values(l2)
	fmt.Printf("l1v : %v\n", l1v)
	fmt.Printf("l2v : %v\n", l2v)
	ret := mergeTwoLists(&l1, &l2, &merge)
	retv := values(*ret)
	l1v = values(l1)
	l2v = values(l2)
	fmt.Printf("retv : %v\n", retv)

	fmt.Printf("l1v : %v\n", l1v)
	fmt.Printf("l2v : %v\n", l2v)

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

func values(list ListNode) []int {
	ret := []int{}
	for {
		ret = append(ret, list.Val)
		if list.Next == nil {
			break
		}
		list = *list.Next
	}
	return ret
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode, merge *ListNode) *ListNode {

	if l1 == nil {
		merge.Next = l2
		return merge.Next
	}
	if l2 == nil {
		merge.Next = l1
		return merge.Next
	}
	if l1.Val < l2.Val {
		merge.Next = l1
		mergeTwoLists(l1.Next, l2, merge.Next) //1
		return merge.Next                      //2
	}
	merge.Next = l2
	mergeTwoLists(l1, l2.Next, merge.Next)
	return merge.Next
}
