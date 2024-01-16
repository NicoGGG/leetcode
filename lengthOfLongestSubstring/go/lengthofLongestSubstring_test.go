package main

import (
	"testing"
)

func TestLengthOfLongestSubstring1(t *testing.T) {
	t1 := "abcabcbb"
	e1 := 3
	r1 := lengthOfLongestSubstring(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestLengthOfLongestSubstring2(t *testing.T) {
	t1 := "bbbbb"
	e1 := 1
	r1 := lengthOfLongestSubstring(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestLengthOfLongestSubstring3(t *testing.T) {
	t1 := "pwwkew"
	e1 := 3
	r1 := lengthOfLongestSubstring(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestLengthOfLongestSubstring4(t *testing.T) {
	t1 := ""
	e1 := 0
	r1 := lengthOfLongestSubstring(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestLengthOfLongestSubstring5(t *testing.T) {
	t1 := " "
	e1 := 1
	r1 := lengthOfLongestSubstring(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestLengthOfLongestSubstring6(t *testing.T) {
	t1 := "dvdf"
	e1 := 3
	r1 := lengthOfLongestSubstring(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestLengthOfLongestSubstring7(t *testing.T) {
	t1 := "abcd"
	e1 := 4
	r1 := lengthOfLongestSubstring(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestLengthOfLongestSubstring8(t *testing.T) {
	t1 := "abcddefghi"
	e1 := 6
	r1 := lengthOfLongestSubstring(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestLengthOfLongestSubstring9(t *testing.T) {
	t1 := "tmmzuxt"
	e1 := 5
	r1 := lengthOfLongestSubstring(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestLengthOfLongestSubstring10(t *testing.T) {
	t1 := "abcdd"
	e1 := 4
	r1 := lengthOfLongestSubstring(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}