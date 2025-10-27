package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRepeatWithStringBuilder(t *testing.T) {
	got := RepeatWithStringBuilder("a")
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRepeatNTimes(t *testing.T) {
	t.Run("0 times", func(*testing.T) {
		got, err := RepeatNTimes("a", 0)

		if err != nil {
			t.Errorf("not expecting error, got %q", err)
		}

		want := ""
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("3 times", func(*testing.T) {
		got, err := RepeatNTimes("b", 3)

		if err != nil {
			t.Errorf("not expecting error, got %q", err)
		}

		want := "bbb"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("negative times", func(t *testing.T) {
		_, err := RepeatNTimes("c", -1)

		if err == nil {
			t.Errorf("expected error, got <nil>")
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}

func BenchmarkRepeatWithStringBuilder(b *testing.B) {
	for b.Loop() {
		RepeatWithStringBuilder("a")
	}
}

func ExampleRepeatNTimes() {
	suffix, err := RepeatNTimes("h", 10)
	if err != nil {
		return
	}

	scream := "A" + suffix + "!"
	fmt.Println(scream)
	// Output: Ahhhhhhhhhh!
}
