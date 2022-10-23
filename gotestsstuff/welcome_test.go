package main

import "testing"

func TestWelcome(t *testing.T){
	got := Welcome()
	want := "Hello, world!"
	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}