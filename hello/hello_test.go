package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("string is not empty", func(t *testing.T) {
		got := Hello("Reza")
		want := "Hello, Reza"

		assertCorrectMessage(t, got, want)
	})
	t.Run("string is empty", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
