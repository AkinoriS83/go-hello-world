package main

import "testing"

func TestMakeHelloWorld(t *testing.T) {
	want := "Hello World! Testing"
	got := makeHelloWorld("Testing")
	if got != want {
		t.Errorf("got = %s; want %s", got, want)
	}
}
