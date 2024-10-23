package main

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	// Create a linked list with a cycle
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node1 // create cycle

	if hasCycle(node1) != true {
		t.Errorf("hasCycle was incorrect, got: false, want: true.")
	}

	// Create a linked list without a cycle
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node6 := &ListNode{Val: 6}
	node4.Next = node5
	node5.Next = node6

	if hasCycle(node4) != false {
		t.Errorf("hasCycle was incorrect, got: true, want: false.")
	}
}

func TestAddTwoNumbers(t *testing.T) {
	// Create two linked lists of equal length
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 4}
	node3 := &ListNode{Val: 3}
	node1.Next = node2
	node2.Next = node3

	node4 := &ListNode{Val: 5}
	node5 := &ListNode{Val: 6}
	node6 := &ListNode{Val: 4}
	node4.Next = node5
	node5.Next = node6

	result := addTwoNumbers(node1, node4)
	if result.Val != 7 || result.Next.Val != 0 || result.Next.Next.Val != 8 {
		t.Errorf("addTwoNumbers was incorrect, got: %d%d%d, want: 708.", result.Val, result.Next.Val, result.Next.Next.Val)
	}

	// Create two linked lists of unequal length
	node7 := &ListNode{Val: 1}
	node8 := &ListNode{Val: 8}
	node7.Next = node8

	node9 := &ListNode{Val: 0}

	result = addTwoNumbers(node7, node9)
	if result.Val != 1 || result.Next.Val != 8 {
		t.Errorf("addTwoNumbers was incorrect, got: %d%d, want: 18.", result.Val, result.Next.Val)
	}
}

func TestMergeTwoLists(t *testing.T) {
	// Create two linked lists of equal length
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 4}
	node1.Next = node2
	node2.Next = node3

	node4 := &ListNode{Val: 1}
	node5 := &ListNode{Val: 3}
	node6 := &ListNode{Val: 4}
	node4.Next = node5
	node5.Next = node6

	result := mergeTwoLists(node1, node4)
	if result.Val != 1 || result.Next.Val != 1 || result.Next.Next.Val != 2 || result.Next.Next.Next.Val != 3 || result.Next.Next.Next.Next.Val != 4 || result.Next.Next.Next.Next.Next.Val != 4 {
		t.Errorf("mergeTwoLists was incorrect, got: %d%d%d%d%d%d, want: 112344.", result.Val, result.Next.Val, result.Next.Next.Val, result.Next.Next.Next.Val, result.Next.Next.Next.Next.Val, result.Next.Next.Next.Next.Next.Val)
	}

	// Create two linked lists of unequal length
	node7 := &ListNode{Val: 1}
	node8 := &ListNode{Val: 2}
	node7.Next = node8

	node9 := &ListNode{Val: 3}
	node10 := &ListNode{Val: 4}
	node11 := &ListNode{Val: 5}
	node9.Next = node10
	node10.Next = node11

	result = mergeTwoLists(node7, node9)
	if result.Val != 1 || result.Next.Val != 2 || result.Next.Next.Val != 3 || result.Next.Next.Next.Val != 4 || result.Next.Next.Next.Next.Val != 5 {
		t.Errorf("mergeTwoLists was incorrect, got: %d%d%d%d%d, want: 12345.", result.Val, result.Next.Val, result.Next.Next.Val, result.Next.Next.Next.Val, result.Next.Next.Next.Next.Val)
	}
}
