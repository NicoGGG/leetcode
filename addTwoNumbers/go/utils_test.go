package addtwonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtils1(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: nil}
	l1.addNumberToList(2)
	l1.addNumberToList(3)
	assert.Equal(t, 1, l1.Val)
	l11 := l1.Next
	assert.Equal(t, 2, l11.Val)
	l12 := l11.Next
	assert.Equal(t, 3, l12.Val)
}

func TestUtils2(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: nil}
	l1.addNumberToList(2)
	assert.Equal(t, 2, listLength(l1))
	l1.addNumberToList(3)
	assert.Equal(t, 3, listLength(l1))
}

func TestListToInt(t *testing.T) {
	l1 := &ListNode{Val: 2}
	l1.addNumberToList(7)

	e := 72
	r := listToInt(l1)
	assert.Equal(t, e, r)
}
