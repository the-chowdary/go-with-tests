package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("asserts if repeated string is equal to wanted count", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		assertCorrectOupt(t, got, want)
	})

	t.Run("asserts if original string is returned when count is passed as 0", func(t *testing.T) {
		got := Repeat("a", 0)
		want := "a"

		assertCorrectOupt(t, got, want)
	})

	t.Run("asserts if original string is returned when count is passed as negative number", func(t *testing.T) {
		got := Repeat("a", -4)
		want := "a"

		assertCorrectOupt(t, got, want)
	})
}

func BechmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func assertCorrectOupt(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("want %q, but got %q", want, got)
	}
}
