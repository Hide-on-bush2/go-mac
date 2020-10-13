package String

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	t.Run("Compare Test", func(t *testing.T) {
		expected := -1
		got := strings.Compare("aaa", "bbb")
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Contains Test", func(t *testing.T) {
		expected := true
		got := strings.Contains("SKT-Faker", "Faker")
		if got != expected {
			t.Errorf("expected '%t' but got '%t'", expected, got)
		}
	})

	t.Run("Count Test", func(t *testing.T) {
		expected := 5
		got := strings.Count("aaaaa", "a")
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}
