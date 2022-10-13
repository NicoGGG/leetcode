package addtwonumbers

import "strconv"

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

var value int
var remains int

func addTwoUnits(l1 *ListNode, l2 *ListNode, remain int) (int, int) {
	s := l1.Val + l2.Val + remain
	r1 := s / 10
	r2 := s % 10
	return r1, r2
}

func (l *ListNode) addNumberToList(i int) {
	next := &ListNode{Val: i, Next: nil}
	for l.Next != nil {
		l = l.Next
	}
	l.Next = next
}

func addNumberToList(l *ListNode, i int) {
	next := &ListNode{Val: i, Next: nil}
	for l.Next != nil {
		l = l.Next
	}
	l.Next = next
}

func appendList(source *ListNode, toAppend *ListNode) {
	var s *ListNode
	for source.Next != nil {
		s = source
		source = source.Next
	}
	for toAppend != nil {
		s.Next = toAppend
		toAppend = toAppend.Next
		s = s.Next
	}
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func listToInt(l *ListNode) int {
	s := ""
	for l != nil {
		s += strconv.Itoa(l.Val)
		l = l.Next
	}
	s = reverse(s)
	r, _ := strconv.Atoi(s)
	return r
}

func intToLisT(i int) *ListNode {
	s := strconv.Itoa(i)
	s = reverse(s)
	r := &ListNode{}
	tmp := r
	for _, c := range s {
		v, _ := strconv.Atoi(string(c))
		tmp.Val = v
		next := &ListNode{}
		tmp.Next = next
	}
	return r
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	value, remains = addTwoUnits(l1, l2, 0)
	lr := &ListNode{Val: remains, Next: nil}
	l1 = l1.Next
	l2 = l2.Next
	for l1 != nil {
		value, remains = addTwoUnits(l1, l2, value)
		addNumberToList(lr, remains)
		l1 = l1.Next
		if l2.Next == nil {
			break
		}
		l2 = l2.Next
	}
	if l1 != nil {
		appendList(lr, l1)
	}
	if l2 != nil {
		appendList(lr, l2)
	}
	return lr
}
