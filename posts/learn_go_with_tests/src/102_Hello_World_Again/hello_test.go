package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("with name", func(t *testing.T) {
		got := Hello("Jacob")
		want := "Hello, Jacob!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("without name", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
