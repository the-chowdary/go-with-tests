package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("asserts if sum of elements in array of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertHelper(t, got, want, numbers)
	})
}

func assertHelper(t testing.TB, got, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d, given: %v", got, want, numbers)
	}
}
