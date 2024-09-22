package iteration

import (
	"fmt"
	"testing"
)

// go test -cover para correr tests con coverage
func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	result := Repeat("1", 5)
	fmt.Println(result)
	// Output: 11111
}
