package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Check name", func(t *testing.T) {
		got := Hello("Nils", "")
		want := "Hello Nils"

		assertCorrectGreeting(t, got, want)
	})

	t.Run("DEFF01", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertCorrectGreeting(t, got, want)
	})

	t.Run("Spanish Test", func(t *testing.T) {
		got := Hello("Teresa", "Spanish")
		want := "Hola Teresa"

		assertCorrectGreeting(t, got, want)
	})

	t.Run("French Test", func(t *testing.T) {
		got := Hello("Josephine", "French")
		want := "Bonjour Josephine"

		assertCorrectGreeting(t, got, want)
	})
}

func assertCorrectGreeting(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
