package add_two_numbers

import "testing"

func Atoll(arr []int) *ListNode {
	first := &ListNode{arr[0], nil}
	previous := first
	for _, n := range arr[1:] {
		previous.Next = &ListNode{n, nil}
		previous = previous.Next
	}

	return first
}

func Lltoa(l *ListNode) []int {
	var res []int
	for l.Next != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return append(res, l.Val)
}

func AssertLinkedList(l1 *ListNode, l2 *ListNode) bool {
	for l1.Next != nil || l2.Next != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	return l1.Next == nil && l2.Next == nil
}

func TestAddTwoNumber(t *testing.T) {
	functions := []struct {
		Name     string
		Function func(*ListNode, *ListNode) *ListNode
	}{
		{"naive", Naive},
		{"optimal", Optimal},
	}

	cases := []struct {
		l1       []int
		l2       []int
		Expected []int
	}{
		{
			[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8},
		},
		{
			[]int{0}, []int{0}, []int{0},
		},
		{
			[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, f := range functions {
		for _, tt := range cases {
			got := f.Function(Atoll(tt.l1), Atoll(tt.l2))
			if !AssertLinkedList(Atoll(tt.Expected), got) {
				t.Errorf("func: %s, l1: %v, l2: %v, expected: %v, got: %v", f.Name, tt.l1, tt.l2, tt.Expected, Lltoa(got))
			}
		}
	}
}
