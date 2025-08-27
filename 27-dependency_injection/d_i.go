package main

import (
	"fmt"
	"io"
)

func Greet(w io.Writer, s string) {
	fmt.Fprint(w, "Hello, ", s)

}

func Greet2(w io.Writer, s string) {
	chunk := []byte{'H', 'e', 'l', 'l', 'o', ',', ' '}
	for i := range s {
		chunk = append(chunk, s[i])
	}

	w.Write(chunk)

}
