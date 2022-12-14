package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGreet(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Greet("Chris", "")
		want := "Hello, Chris"

		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }
		assertCorrectMessage(t, got, want)

	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Greet("", "")
		want := "Hello, World"

		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }
		assertCorrectMessage(t, got, want)
	})
	// Spanish
	t.Run("in Spanish", func(t *testing.T) {
		got := Greet("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	// French
	t.Run("in French", func(t *testing.T) {
		got := Greet("Monde", "French")
		want := "Bonjour, Monde"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
