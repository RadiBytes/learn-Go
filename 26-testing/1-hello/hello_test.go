package main

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleHello() {
	fmt.Println(Hello())
	// Output: Hello, testing ground
}

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, testing ground"

	assertCorrectMessage(t, got, want)
}

func TestHelloWithArg(t *testing.T) {
	got := Hello2("Ra")
	want := "Hello, testing ground...Ra"
	got2 := Hello2("")
	want2 := "Hello, testing ground...People"

	gotKind := reflect.ValueOf(want).Kind()
	reqKind := reflect.String

	if gotKind != reqKind {
		t.Errorf("Req kind: %s, got: %s", reqKind, gotKind)
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
		assertCorrectMessage(t, got2, want2)
	}
}
func TestHelloWithArgGrouped(t *testing.T) {
	t.Run("Test with Ra", func(t *testing.T) {
		got := Hello2("Ra")
		want := "Hello, testing ground...Ra"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Test with empty string and type subtest", func(t *testing.T) {
		got := Hello2("")
		want := "Hello, testing ground...People"
		assertCorrectMessage(t, got, want)

		t.Run("Type test", func(t *testing.T) {
			gotKind := reflect.ValueOf(want).Kind()
			reqKind := reflect.String
			if gotKind != reqKind {
				t.Errorf("Req kind: %s, got: %s", reqKind, gotKind)
			}

		})
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
