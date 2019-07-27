package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArraySum(t *testing.T) {
	t.Run("collection of any size of numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := ArraySum(numbers)
		expected := 15
		if got != expected {
			t.Errorf("expected %d but got %d given %v", expected, got, numbers)
		}
	})
}

func ExampleArraySum() {
	numbers := []int{1, 2, 3, 4, 5}
	res := ArraySum(numbers)
	fmt.Println(res)
	// Output:15
}

func BenchmarkArraySum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		ArraySum(numbers)
	}
}

func TestSlicesSum(t *testing.T) {
	t.Run("sum slices", func(t *testing.T) {
		slices := [][]int{[]int{1, 2}, []int{3, 4}, []int{5}}
		got := SlicesSum(slices)
		expected := []int{3, 7, 5}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %d but got %d given %v", expected, got, slices)
		}
	})
}

func ExampleSlicesSum() {
	slices := [][]int{[]int{1, 2}, []int{3, 4}, []int{9, 2}}
	res := SlicesSum(slices)
	fmt.Println(res)
	// Output:[3 7 11]
}

func BenchmarkSlicesSum(b *testing.B) {
	slices := [][]int{[]int{1, 2}, []int{3, 4}, []int{9, 2}}
	for i := 0; i < b.N; i++ {
		SlicesSum(slices)
	}
}
