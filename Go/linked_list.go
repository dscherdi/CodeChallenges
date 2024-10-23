package main

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

// https://leetcode.com/problems/add-two-numbers/?envType=study-plan-v2&envId=top-interview-150
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode = nil
	previousElement := result
	head := result
	carry := 0

	for l1.Next != nil && l2.Next != nil {

		sum := l1.Val + l2.Val + carry
		carry = sum / 10

		newElement := &ListNode{Val: sum % 10}
		previousElement = head
		head = newElement
		previousElement.Next = head

		l1 = l1.Next
		l2 = l2.Next
	}

	return result
}

// https://leetcode.com/problems/merge-two-sorted-lists/?envType=study-plan-v2&envId=top-interview-150
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result *ListNode

	var previous *ListNode
	var head *ListNode

	for list1 != nil || list2 != nil {

		previous = head

		if list1 == nil || (list2 != nil && list2.Val <= list1.Val) {
			head = list2
			list2 = list2.Next
		} else if list2 == nil || list1.Val < list2.Val {
			head = list1
			list1 = list1.Next
		}

		if previous != nil {
			previous.Next = head
		}

		if result == nil {
			result = head
		}
	}

	return result
}
