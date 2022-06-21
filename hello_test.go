package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to a person", func(t *testing.T) {
		got := Hello("Brett", "")
		want := "Hello, Brett"
		assertCorrectMessage(t, got, want)

	})

	t.Run("say Hello, World when string supplied is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Johnny", "French")
		want := "Bonjour, Johnny"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Italian", func(t *testing.T) {
		got := Hello("Vanessa", "Italian")
		want := "Ciao, Vanessa"
		assertCorrectMessage(t, got, want)
	})
}