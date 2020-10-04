package atoi

func myAtoi(s string) int {
	r := 0
	n := 1
	p := false
	for _, c := range s {
		if c == 32 {
			if p {
				break
			}
		} else if c < 48 || c > 57 {
			if (c == 43 || c == 45) && p == false {
				p = true
				if c == 45 {
					n = -1
				}
			} else {
				break
			}
		} else {
			p = true
			r = r * 10 + int(c) - 48
		}
		if r * n >= (1<<31) {
			return (1<<31 - 1)
		} else if r * n < (-1<<31) {
			return (-1<<31)
		}
	}
	return r * n
}