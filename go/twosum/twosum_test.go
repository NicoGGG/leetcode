package main

import (
	"reflect"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	t1 := []int{2, 7, 11, 15}
	e1 := []int{0, 1}
	r1 := twoSum(t1, 9)
	if !reflect.DeepEqual(r1, e1) {
		t.Errorf("Expected %v", e1)
	}
}

func TestTwoSum2(t *testing.T) {
	t2 := []int{3, 3}
	e2 := []int{0, 1}
	r2 := twoSum(t2, 6)
	if !reflect.DeepEqual(r2, e2) {
		t.Errorf("Expected %v ; Got %v", e2, r2)
	}
}