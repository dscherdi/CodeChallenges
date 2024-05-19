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
