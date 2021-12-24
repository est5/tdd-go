package iteration

import (
	"fmt"
	"testing"
)

const repeatTimes = 2

func TestRepeat(t *testing.T) {
	result := Repeat("a", repeatTimes)
	expected := "aa"

	if result != expected {
		t.Errorf("expected %q got %q", expected, result)
	}

}

func ExampleRepeat() {
	repeated := Repeat("abc", 3)
	fmt.Println(repeated)
	//Output: abcabcabc

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}
