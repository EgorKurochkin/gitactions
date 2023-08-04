package main

import "testing"

func TestHello(t *testing.T) {
	name := "Bob"
	want := "Hello Bob"

	if got := Hello(name); got != want {
		t.Errorf("hello() = %q, want %q", got, want)
	}
}
