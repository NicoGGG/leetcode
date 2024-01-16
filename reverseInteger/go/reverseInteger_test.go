package main

import (
"testing"
)

func TestReverse1(t *testing.T) {
	i := 123
	e := 321
	r := reverse(i)
	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

func TestReverse2(t *testing.T) {
	i := -123
	e := -321
	r := reverse(i)
	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

func TestReverse3(t *testing.T) {
	i := 120
	e := 21
	r := reverse(i)
	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

func TestReverse4(t *testing.T) {
	i := 0
	e := 0
	r := reverse(i)
	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

func TestReverse5(t *testing.T) {
	i := 1534236469
	e := 0
	r := reverse(i)
	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}