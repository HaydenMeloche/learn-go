package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: %q wanted: %q", got, want)
		}
	}

	t.Run("saying hello to myself", func(t *testing.T) {
		got := Hello("Hayden")
		want := "Hello Hayden"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when no name is provided", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})

}
