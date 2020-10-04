package atoi

import (
	"testing"
)

func TestMyAtoi1(t *testing.T) {
	t1 := "42"
	e1 := 42
	r1 := myAtoi(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestMyAtoi2(t *testing.T) {
	t1 := "   -42"
	e1 := -42
	r1 := myAtoi(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestMyAtoi3(t *testing.T) {
	t1 := "4193 with words"
	e1 := 4193
	r1 := myAtoi(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestMyAtoi4(t *testing.T) {
	t1 := "words and 987"
	e1 := 0
	r1 := myAtoi(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestMyAtoi5(t *testing.T) {
	t1 := "-91283472332"
	e1 := -2147483648
	r1 := myAtoi(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestMyAtoi6(t *testing.T) {
	t1 := "21474836460"
	e1 := 2147483647
	r1 := myAtoi(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestMyAtoi7(t *testing.T) {
	t1 := "00000-42a1234"
	e1 := 0
	r1 := myAtoi(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestMyAtoi8(t *testing.T) {
	t1 := "2147483648"
	e1 := 2147483647
	r1 := myAtoi(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestMyAtoi9(t *testing.T) {
	t1 := "9223372036854775808"
	e1 := 2147483647
	r1 := myAtoi(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}
