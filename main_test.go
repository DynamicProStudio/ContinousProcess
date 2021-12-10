package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, DynamicPro Studio"
	got := hello("DynamicPro Studio")

	if want != got {
		t.Fatalf("want %s, but got %s\n", want, got)
	}
}

func TestHelloFail(t *testing.T) {
	want := "Ada"
	got := hello("Peter")

	if want == got {
		t.Fatalf("want %s, but got %s", want, got)
	}
}
