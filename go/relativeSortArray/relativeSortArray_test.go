package relativeSortArray

import (
	"reflect"
	"testing"
)

func TestRelativeSortArray1(t *testing.T) {
	arr1 := []int{2,3,1,3,2,4,6,7,9,2,19}
	arr2 := []int{2,1,4,3,9,6}
	e := []int{2,2,2,1,4,3,3,9,6,7,19}
	r := relativeSortArray(arr1, arr2)
	if !reflect.DeepEqual(r, e) {
		t.Errorf("Expected %v, got %v", e, r)
	}
}
