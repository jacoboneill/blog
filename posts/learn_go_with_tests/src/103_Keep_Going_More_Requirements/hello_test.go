package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("no name", func(t *testing.T) {
		got := Hello("", "en")
		want := "Hello, World!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("with name", func(t *testing.T) {
		got := Hello("Jacob", "en")
		want := "Hello, Jacob!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("in English", func(t *testing.T) {
		got := Hello("Jacob", "en")
		want := "Hello, Jacob!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("default to English", func(t *testing.T) {
		got := Hello("Jacob", " ")
		want := "Hello, Jacob!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Jacob", "es")
		want := "Hola, Jacob!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Jacob", "fr")
		want := "Bonjour, Jacob!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
