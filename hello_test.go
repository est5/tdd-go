package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Ilya")
	want := "Hello, Ilya"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
