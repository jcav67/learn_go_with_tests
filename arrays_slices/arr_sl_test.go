package arrays_slices

import (
	"slices"
	"testing"
)

// func TestSum(t *testing.T) {

// 	numbers := [5]int{1, 2, 3, 4, 5}

// 	got := Sum(numbers)
// 	want := 15

// 	if got != want {
// 		t.Errorf("got %d want %d given, %v", got, want, numbers)
// 	}
// }

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll_longway([]int{1, 2, 3}, []int{3, 4, 5})
	want := []int{6, 12}

	if !slices.Equal(got, want) {
		t.Errorf("got %d want %d given", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %d want %d given", got, want)
		}
	}
	t.Run("slices with numbers", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{3, 4, 5})
		want := []int{5, 9}

		checkSums(t, got, want)
	})

	t.Run("at least one empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
