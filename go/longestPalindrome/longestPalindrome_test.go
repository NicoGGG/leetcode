package longestpalindrome

import (
	"log"
	"testing"
)

func TestIsPalindrome1(t *testing.T) {
	i := "aba"
	e := true
	r := isPalindrome(i)
	if e != r {
		log.Fatalf("Expected result: %v, got %v", e, r)
	}
}

func TestIsPalindrome2(t *testing.T) {
	i := "ab"
	e := false
	r := isPalindrome(i)
	if e != r {
		log.Fatalf("Expected result: %v, got %v", e, r)
	}
}

func TestIsPalindrome3(t *testing.T) {
	i := "bb"
	e := true
	r := isPalindrome(i)
	if e != r {
		log.Fatalf("Expected result: %v, got %v", e, r)
	}
}

func TestIsPalindrome4(t *testing.T) {
	i := ""
	e := true
	r := isPalindrome(i)
	if e != r {
		log.Fatalf("Expected result: %v, got %v", e, r)
	}
}

func TestIsPalindrome5(t *testing.T) {
	i := "a"
	e := true
	r := isPalindrome(i)
	if e != r {
		log.Fatalf("Expected result: %v, got %v", e, r)
	}
}

func TestLongestPalindrome1(t *testing.T) {
	s := "babad"
	r := longestPalindrome(s)
	e := "bab"
	if r != e {
		log.Fatalf("Expected result: %s, got %s", e, r)
	}
}

func TestLongestPalindrome2(t *testing.T) {
	s := "cbbd"
	r := longestPalindrome(s)
	e := "bb"
	if r != e {
		log.Fatalf("Expected result: %s, got %s", e, r)
	}
}

func TestLongestPalindrome3(t *testing.T) {
	s := "a"
	r := longestPalindrome(s)
	e := "a"
	if r != e {
		log.Fatalf("Expected result: %s, got %s", e, r)
	}
}

func TestLongestPalindrome4(t *testing.T) {
	s := "acaaaaaaaaa"
	r := longestPalindrome(s)
	e := "aaaaaaaaa"
	if r != e {
		log.Fatalf("Expected result: %s, got %s", e, r)
	}
}
