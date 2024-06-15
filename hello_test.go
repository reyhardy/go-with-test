package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Reza")
	want := "Hello, Reza"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
