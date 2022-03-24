package longestpalindrome

func longestPalindrome(s string) string {
	return ""
}

func isPalindrome(s string) bool {
	l := len(s)
	for i, c := range s {
		if byte(c) != s[l-1-i] {
			return false
		}
	}
	return true
}
