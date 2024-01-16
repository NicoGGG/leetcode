package longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	r := ""
	if len(strs) == 0 {
		return r
	}
	l := len(strs[0])
	var c uint8
	m := make(map[uint8]bool)
	for i := 0; i < l; i++ {
		for _, s := range strs {
			if len(s) <= i {
				return r
			}
			c = s[i]
			m[c] = true
		}
		if len(m) != 1 {
			return r
		}
		m = make(map[uint8]bool)
		r = r + string(c)
	}
	return r
}