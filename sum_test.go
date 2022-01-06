package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("The Sum function shall add all elements of an array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given: %v", got, want, numbers)
		}
	})

	t.Run("The Sum function shall account for an empty array", func(t *testing.T) {
		numbers := []int{}
		got := Sum(numbers)
		want := 0
		if got != want {
			t.Errorf("got %d want %d given: %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("The SumAll function shall sum a varying number of array slices and return a slice with the totals for arch slice", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("The SumAllTails function shall sum the tail of each array slice passed", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("The SumAllTails function shall account for empty arrays", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
