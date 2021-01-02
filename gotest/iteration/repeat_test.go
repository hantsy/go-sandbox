package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "AAAAA"

	if repeated != expected {
		t.Errorf("exptected %q but go %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	Repeat("a", b.N)
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Printf(repeated)
	// Output: AAAAA
}
