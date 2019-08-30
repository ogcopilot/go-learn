package iteration

import (
	"math/rand"
	"testing"
	"time"
)

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Repeat 'x' {n} of times for n = 3", func(t *testing.T) {
		repeated, _ := Repeat("x", 3)
		expected := "xxx"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("Repeat 'x' {n} of times for n = 0", func(t *testing.T) {
		repeated, _ := Repeat("x", 0)
		expected := ""
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("Repeat 'y' {n} of times for n = 1", func(t *testing.T) {
		repeated, _ := Repeat("y", 1)
		expected := "y"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("Repeat 'x' {n} of times for n = -1", func(t *testing.T) {
		_, err := Repeat("x", -1)
		expect := true
		got := err != nil
		if got != expect {
			t.Errorf("got %t wanted %t", got, expect)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		Repeat("a", r.Intn(10))
	}
}
