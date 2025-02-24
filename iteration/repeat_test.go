package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("default test", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("repeated %q, but expected %q", repeated, expected)
		}
	})

	t.Run("Add number of repeat", func(t *testing.T) {
		repeated := Repeat("a", 2)
		expected := "aa"

		if repeated != expected {
			t.Errorf("repeated: %q,but expected: %q", repeated, expected)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 2)
	}
}

func ExampleRepeat() {
	repeating := Repeat("a", 3)

	fmt.Println(repeating) // Output: aaa
}
