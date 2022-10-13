package addtwonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertListEqual(t *testing.T, expected *ListNode, actual *ListNode) {
	assert.NotNil(t, actual)
	assert.Equal(t, listLength(expected), listLength(actual), "List length doesn't match")
	for expected != nil {
		assert.Equal(t, expected.Val, actual.Val)
		expected = expected.Next
		actual = actual.Next
	}
}

func TestAddTwoUnits(t *testing.T) {
	l1 := &ListNode{Val: 2}
	l1.addNumberToList(7)

	l2 := &ListNode{Val: 5}
	l2.addNumberToList(6)

	e1, e2 := 0, 7
	r1, r2 := addTwoUnits(l1, l2, 0)
	assert.Equal(t, e1, r1)
	assert.Equal(t, e2, r2)

	l1 = l1.Next
	l2 = l2.Next
	e11, e21 := 1, 6
	r11, r21 := addTwoUnits(l1, l2, r2)
	assert.Equal(t, e11, r11)
	assert.Equal(t, e21, r21)
}

func TestIntToList(t *testing.T) {
	i := 456
	r := intToLisT(i)

	e := &ListNode{Val: 6}
	e.addNumberToList(5)
	e.addNumberToList(4)

	assertListEqual(t, e, r)
}

func TestExample1(t *testing.T) {
	l1 := &ListNode{Val: 2}
	l1.addNumberToList(4)
	l1.addNumberToList(3)

	l2 := &ListNode{Val: 5}
	l2.addNumberToList(6)
	l2.addNumberToList(4)

	le := &ListNode{Val: 7}
	le.addNumberToList(0)
	le.addNumberToList(8)

	lr := addTwoNumbers(l1, l2)
	assertListEqual(t, le, lr)

}

func TestExample2(t *testing.T) {
	l1 := &ListNode{Val: 0}

	l2 := &ListNode{Val: 0}

	le := &ListNode{Val: 0}

	lr := addTwoNumbers(l1, l2)
	assertListEqual(t, le, lr)
}

func TestExample3(t *testing.T) {
	l1 := &ListNode{Val: 9}
	l1.addNumberToList(9)
	l1.addNumberToList(9)
	l1.addNumberToList(9)
	l1.addNumberToList(9)
	l1.addNumberToList(9)
	l1.addNumberToList(9)

	l2 := &ListNode{Val: 9}
	l2.addNumberToList(9)
	l2.addNumberToList(9)
	l2.addNumberToList(9)

	le := &ListNode{Val: 8}
	le.addNumberToList(9)
	le.addNumberToList(9)
	le.addNumberToList(9)
	le.addNumberToList(0)
	le.addNumberToList(0)
	le.addNumberToList(0)
	le.addNumberToList(1)

	lr := addTwoNumbers(l1, l2)
	assertListEqual(t, le, lr)

}
