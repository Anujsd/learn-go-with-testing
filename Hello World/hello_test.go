package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say Hello to people", func(t *testing.T) {
		got := hello("Anuj")
		want := "Hello Anuj"
		assertMessage(t, got, want)
	})

	t.Run("Say Hello World when empty", func(t *testing.T) {
		got := hello("")
		want := "Hello World"
		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q Want %q", got, want)
	}
}
