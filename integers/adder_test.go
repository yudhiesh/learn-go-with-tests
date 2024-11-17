package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(1, 2)
	expected := 3
	if sum != expected {
		t.Errorf("got '%d', want '%d'", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
}
