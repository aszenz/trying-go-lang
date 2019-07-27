package arrays

import (
	"fmt"
	"testing"
)

func ExampleArraySum() {
	numbers := []int{1, 2, 3, 4, 5}
	res := ArraySum(numbers)
	fmt.Println(res)
	// Output:15
}

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

func BenchmarkArraySum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		ArraySum(numbers)
	}
}
