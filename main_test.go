package main

import "testing"

func TestSayHello(t *testing.T) {

	got := sayHello()
	want := "Hello Golang!"

	if got != want {
		t.Fatalf("want %s got %s", want, got)
	}
}
