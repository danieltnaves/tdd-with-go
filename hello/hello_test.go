package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("World", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello with an empty string supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in portuguese from brazil", func(t *testing.T) {
		got := Hello("Mundo", "pt-BR")
		want := "Olá, Mundo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in german", func(t *testing.T) {
		got := Hello("Welt", "de")
		want := "Hallo, Welt"
		assertCorrectMessage(t, got, want)
	})

}
