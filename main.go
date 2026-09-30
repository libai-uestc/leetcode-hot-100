package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)

	head.Next.Next = head

	head.Next = nil

	return newHead
}

// func reverseList(head *ListNode) *ListNode {
// 	var prev *ListNode
// 	curr := head

// 	for curr != nil {
// 		nextTemp := curr.Next
// 		curr.Next = prev
// 		prev = curr
// 		curr = nextTemp
// 	}
// 	return prev
// }
