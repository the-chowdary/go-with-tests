package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to person", func(t *testing.T) {
		got := Hello("The Chowdary")
		want := "Hello, The Chowdary"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, world!' when an empty string is passed", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
