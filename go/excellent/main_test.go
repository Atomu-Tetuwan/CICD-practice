package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(2)
	if result != "Even" {
		t.Errorf("Expected Even, but got %s", result)
	}
}
