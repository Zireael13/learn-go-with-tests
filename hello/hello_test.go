package main

import "testing"

func TestHello(t *testing.T) {

	assertString := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("with name param", func(t *testing.T) {
		got := Hello("Matt")
		want := "Hello, Matt"

		assertString(t, got, want)
	})

	t.Run("with empty name param", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertString(t, got, want)
	})
}
