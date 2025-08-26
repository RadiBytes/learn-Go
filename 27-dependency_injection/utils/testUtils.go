package utils

import (
	"testing"
)

func AssertCorrectMessage(t testing.TB, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("\ngot: %v\nwant: %v\n", got, want)
	}
}
