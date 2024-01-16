package longestpalindrome

import (
	"testing"
)

func BenchmarkLongestPalindrome(b *testing.B) {
	s := "babad"
	for n := 0; n < b.N; n++ {
		longestPalindrome(s)
	}
}
