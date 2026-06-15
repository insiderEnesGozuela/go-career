package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// bytes.Buffer Writer interface'ini implement ediyor (Write metodu var)
	// bu yüzden Writer olarak verip içine ne yazıldığını kontrol edebiliyoruz
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
