package testhelpers

import (
	"reflect"
	"testing"
)

// CheckSlicesSum checks if the slices sum is equal to the wanted array
func CheckSlicesSum(t *testing.T, got []int, want []int, given [][]int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %d but got %d given %v", want, got, given)
	}
}

// CheckArraySum checks if the array sum equals given wanted value
func CheckArraySum(t *testing.T, got int, want int, given []int) {
	t.Helper()
	if got != want {
		t.Errorf("expected %d but got %d given %v", want, got, given)
	}
}
