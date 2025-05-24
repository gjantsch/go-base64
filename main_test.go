package main

import "testing"

func TestEncoder(t *testing.T) {
	want := "YSBzaW1wbGUgdGVzdA=="
	got := Base64Encode([]byte("a simple test"))
	if got != want {
		t.Fatalf("Wanted %s but got %s", want, got)
	}
}

func TestDecoder(t *testing.T) {
	want := "a simple test"
	got := Base64Decode([]byte("YSBzaW1wbGUgdGVzdA=="))
	if got != want {
		t.Errorf("Wanted '%v' but got '%v'", []byte(want), []byte(got))
	}
}
