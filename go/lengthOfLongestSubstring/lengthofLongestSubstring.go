package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[int32]int)
	tR := 0
	r := 0
	b := 0
	for i, c := range s {
		if j, ok := m[c]; ok {
			// If char c is in the map, all substring starting at or before m[c] are ended
			if j+1 > b {
				// New beginning becomes m[c] + 1, unless we are on a char that was before b, because its streak was already ended
				// For example, "abccadefgh" when we get to the second 'c' at index 4, beginning is now 4.
				//  when we get to the second 'a' at index 5, beginning doesn't change because the 'a' starting string was already over
				//  when we met the second 'c' at index 4
				b = j + 1
			}
		}
		// We adjust new longest string by comparing the current streak, to the longest one so far
		tR = i - b + 1
		if tR > r {
			r = tR
		}
		// Whether char is in the map or not, m[c] becomes i
		m[c] = i
	}
	return r
}
