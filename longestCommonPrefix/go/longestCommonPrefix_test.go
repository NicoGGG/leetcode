package longestCommonPrefix

import (
	"testing"
)

func TestLongestCommonPrefix1(t *testing.T) {
	t1 := []string{"flower", "flow", "flight"}
	e1 := "fl"
	r1 := longestCommonPrefix(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestLongestCommonPrefix2(t *testing.T) {
	t1 := []string{"dog", "racecar", "car"}
	e1 := ""
	r1 := longestCommonPrefix(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestLongestCommonPrefix3(t *testing.T) {
	t1 := []string{"", "coucou", "coupbas"}
	e1 := ""
	r1 := longestCommonPrefix(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestLongestCommonPrefix4(t *testing.T) {
	t1 := []string{"coucou", "coucouhaut", "coucoubas"}
	e1 := "coucou"
	r1 := longestCommonPrefix(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestLongestCommonPrefix5(t *testing.T) {
	t1 := []string{}
	e1 := ""
	r1 := longestCommonPrefix(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func TestLongestCommonPrefix6(t *testing.T) {
	t1 := []string{"ab", "a"}
	e1 := "a"
	r1 := longestCommonPrefix(t1)
	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}
}

func BenchmarkLongestCommonPrefix1(b *testing.B) {
	b.ReportAllocs()
	t1 := []string{"flower", "flow", "flight"}
	e1 := "fl"
	for i := 0; i < b.N; i++ {
		r1 := longestCommonPrefix(t1)
		if r1 != e1 {
			b.Errorf("Expected %v, got %v", e1, r1)
		}
	}
}