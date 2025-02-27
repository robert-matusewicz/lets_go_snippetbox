package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, actual T, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("got %v; want %v", actual, expected)
	}
}

func StringContains(t *testing.T, actual string, substr string) {
	t.Helper()

	if !strings.Contains(actual, substr) {
		t.Errorf("got: %q; expected to contain: %q", actual, substr)
	}
}

func NilError(t *testing.T, actual error) {
	t.Helper()

	if actual != nil {
		t.Errorf("got: %v; expected: nil", actual)
	}
}
