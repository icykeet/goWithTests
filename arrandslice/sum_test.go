package arrandslice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		nums := []int{1, 2}

		got := Sum(nums)
		want := 3

		if got != want {
			t.Errorf("got %d, but want %d, given %v", got, want, nums)
		}

	})
}

func TestSumAll(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{1, 2})
		want := []int{6, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %d, got %d", want, got)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("default tailt", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

func BenchmarkTestAllTail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{1, 2}, []int{0, 9})
	}
}
