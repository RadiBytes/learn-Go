package main

import (
	"dependency_injection/utils"

	"testing"
)

type mockItem struct {
	data  []byte
	count int
}

func (m *mockItem) Write(b []byte) (int, error) {
	m.count++
	m.data = b
	return 0, nil
}

func (m mockItem) String() string {
	return string(m.data)
}

func TestMock(t *testing.T) {
	t.Run("Test with Mock", func(t *testing.T) {
		t.Run("Test that Greet can accept mockItem", func(t *testing.T) {
			mockItem := mockItem{}
			Greet(&mockItem, "Ra")
			got := mockItem.String()
			want := "Hello, Ra"
			utils.AssertCorrectMessage(t, got, want)
		})
		t.Run("Test that Greet2 can accept mockItem", func(t *testing.T) {
			mockItem := mockItem{}
			Greet2(&mockItem, "Ra")
			got := mockItem.String()
			want := "Hello, Ra"
			utils.AssertCorrectMessage(t, got, want)
		})
	})

}
