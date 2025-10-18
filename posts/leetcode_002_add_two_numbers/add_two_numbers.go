package add_two_numbers

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Naive(l1 *ListNode, l2 *ListNode) *ListNode {
	traverseListReversed := func(l *ListNode) int {
		i := 0.0
		s := 0
		for l.Next != nil {
			s += int(math.Pow(10, i)) * l.Val
			i++
			l = l.Next
		}

		s += int(math.Pow(10, i)) * l.Val
		return s
	}

	s := traverseListReversed(l1) + traverseListReversed(l2)

	generateLinkedList := func(s int) *ListNode {
		first := ListNode{s % 10, nil}
		previous := &first
		s /= 10

		for s != 0 {
			previous.Next = &ListNode{s % 10, nil}
			previous = previous.Next
			s /= 10
		}

		return &first
	}

	return generateLinkedList(s)
}

func Optimal(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		val := carry

		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}

		carry = val / 10

		curr.Next = &ListNode{val % 10, nil}
		curr = curr.Next
	}

	return dummy.Next
}
