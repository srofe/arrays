package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("The Sum function shall add all elements of an array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given: %v", got, want, numbers)
		}
	})

	t.Run("The Sum function shall add collections of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d given: %v", got, want, numbers)
		}
	})
}
