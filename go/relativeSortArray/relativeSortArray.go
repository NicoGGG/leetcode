package relativeSortArray

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int][]int)
	r := []int{}
	for k, v := range arr2 {
		m1[v] = k
	}
	last := []int{}
	for _, j := range arr1 {
		if i, ok := m1[j]; ok  {
			m2[i] = append(m2[i], j)
		} else {
			last = append(last, j)
			continue
		}
	}
	keys := make([]int, 0, len(m2))
	for k := range m2 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		r = append(r, m2[k]...)
	}
	sort.Ints(last)
	r = append(r, last...)
	return r
}