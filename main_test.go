package main

import "testing"

func TestMakeHelloWorld(t *testing.T) {
	want := "Hello World! by myapp Testing"
	got := makeHelloWorld("Testing")
	if got != want {
		t.Errorf("got = %s; want %s", got, want)
	}
}
