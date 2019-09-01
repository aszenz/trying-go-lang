package integers

import (
	"fmt"
	"testing"

	"github.com/aszenz/basic-golang/go-test-helpers"
)

func ExampleAdd() {
	sum := Add(2, 3)
	fmt.Println(sum)
	// Output: 5
}
func TestAdder(t *testing.T) {
	sum := Add(2, 3)
	expected := 5
	testhelpers.CheckIntegersEqual(t, sum, expected, "add")
}
