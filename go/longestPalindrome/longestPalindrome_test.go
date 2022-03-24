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

func TestLongestPalindrome1(t *testing.T) {
	s := "babad"
	r := longestPalindrome(s)
	e := "bab"
	if r != e {
		log.Fatalf("Expected result: %s, got %s", e, r)
	}
}
