package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Tony", "spanish")
	want := "Hola, Tony"

	t.Run("Test Hello function", func(t *testing.T) {
		if got != want {
			t.Error("Incorrect reference")
		}
	})
}
