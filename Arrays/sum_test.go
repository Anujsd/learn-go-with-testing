package arrays

import "testing"

func TestSum(t *testing.T) {

	t.Run("Sum of 4 Elemetns", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got := GetSum(numbers)
		want := 10
		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})
	t.Run("Sum of 5 Elemetns", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := GetSum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})

}
