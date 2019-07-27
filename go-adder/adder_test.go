package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(2, 3)
	fmt.Println(sum)
	// Output: 5
}
func TestAdder(t *testing.T) {
	sum := Add(2, 3)
	expected := 5
	if sum != expected {
		t.Errorf("expected %d but got '%d'", expected, sum)
	}
}
