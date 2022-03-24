package longestpalindrome

func longestPalindrome(s string) string {
	r := ""
	for i := 0; i < len(s); i++ {
		for j := len(r) + i + 1; j < len(s)+1; j++ {
			if isPalindrome(s[i:j]) {
				r = s[i:j]
			}
		}
	}
	return r
}

func isPalindrome(s string) bool {
	l := len(s)
	for i := range s {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}
