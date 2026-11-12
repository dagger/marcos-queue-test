package main

import "testing"

func TestPass(t *testing.T) {
	if got := 1 + 1; got != 2 {
		t.Fatalf("1+1 = %d, want 2", got)
	}
}
