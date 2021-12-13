package main

import "testing"

func TestAdd(t *testing.T) {
	if v := Add(1, 2); v != 3 {
		t.Fatal("bad add")
	}
}
