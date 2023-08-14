package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	current := dummy

	for current.Next != nil && current.Next.Next != nil {
		first := current.Next
		second := current.Next.Next

		// Swapping
		first.Next = second.Next
		current.Next = second
		current.Next.Next = first

		current = current.Next.Next
	}

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	// Example usage
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	fmt.Println("Original List:")
	printList(head)

	newHead := swapPairs(head)
	fmt.Println("List after Swapping Pairs:")
	printList(newHead)
}
