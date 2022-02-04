package main

import (
	"fmt"
	"testing"
)

func TestMainPrintFirst(t *testing.T) {

	p := Person{First: "Paul"}

	PrintFirst(p)
}

func ExamplePrintFirst() {

	p := Person{First: "Paul"}

	got := PrintFirst(p)
	fmt.Printf("%s", got)
	// Output: Paul
}
