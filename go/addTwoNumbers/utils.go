package addtwonumbers

func listLength(l *ListNode) int {
	r := 0
	for l != nil {
		r++
		l = l.Next
	}
	return r
}
