package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeatCharacter() {
	res := RepeatCharacter("i", 3)
	fmt.Println(res)
	// Output: iii
}
func TestIteration(t *testing.T) {
	res := RepeatCharacter("h", 5)
	expected := "hhhhh"
	if res != expected {
		t.Errorf("expected %s but got '%s'", expected, res)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatCharacter("a", 5)
	}
}
