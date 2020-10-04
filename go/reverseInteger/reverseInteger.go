package main

func reverse(x int) int {
	r := 0
	for ; x != 0 ; x /= 10 {
		r = r * 10 + x % 10
	}
	if r > (1<<31) || r < (-1<<31) {
		return 0
	}
	return r
}
