package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to person in english", func(t *testing.T) {
		got := Hello("The Chowdary", "English")
		want := "Hello, The Chowdary"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, world!' in english when an empty strings for both name and language are passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hola, world!' in spanish", func(t *testing.T) {
		got := Hello("", "spanish")
		want := "Hola, world!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hola, The Chowdary!' in spanish", func(t *testing.T) {
		got := Hello("The Chowdary", "spanish")
		want := "Hola, The Chowdary"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Bonjour, world!' in french", func(t *testing.T) {
		got := Hello("", "french")
		want := "Bonjour, world!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Bonjour, The Chowdary!' in french", func(t *testing.T) {
		got := Hello("The Chowdary", "french")
		want := "Bonjour, The Chowdary"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
