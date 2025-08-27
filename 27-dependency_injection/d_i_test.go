package main

import (
	"bytes"
	"dependency_injection/utils"
	"net/http/httptest"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("Test that Greet can accept Buffer", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Ra")
		got := buffer.String()
		want := "Hello, Ra"
		utils.AssertCorrectMessage(t, got, want)
	})

	t.Run("Test that Greet can accept http writer", func(t *testing.T) {

		w := httptest.NewRecorder()
		// var w http.ResponseWriter
		Greet(w, "Ra")
		got := w.Body.String()
		want := "Hello, Ra"
		utils.AssertCorrectMessage(t, got, want)
	})

}
