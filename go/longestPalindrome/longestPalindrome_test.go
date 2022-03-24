package longestpalindrome

import (
	"log"
	"testing"
)

func TestLongestPalindrome1(t *testing.T) {
	s := "babad"
	r := longestPalindrome(s)
	e := "bab"
	if r != e {
		log.Fatalf("Expected result: %s, got %s", e, r)
	}
}
