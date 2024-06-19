package integer

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		got := Add(3, 2)
		want := 5

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 6)
	fmt.Println(sum)
	// Output: 7
}
