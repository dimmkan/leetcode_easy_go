/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.



Example 1:


Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: list1 = [], list2 = []
Output: []
Example 3:

Input: list1 = [], list2 = [0]
Output: [0]


Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.
*/
/**
 * Definition for singly-linked list.
 *
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	merge := &ListNode{}
	head := merge

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			merge.Next = list1
			list1 = list1.Next
		} else {
			merge.Next = list2
			list2 = list2.Next
		}
		merge = merge.Next
	}

	for list1 != nil {
		merge.Next = list1
		list1 = list1.Next
		merge = merge.Next
	}

	for list2 != nil {
		merge.Next = list2
		list2 = list2.Next
		merge = merge.Next
	}

	return head.Next
}
