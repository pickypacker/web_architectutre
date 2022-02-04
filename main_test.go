package main

import "testing"

func TestMainPrintFirst(t *testing.T) {

	p := Person{First: "Paul"}

	PrintFirst(p)
}

func ExamplePrintFirst() {

	p := Person{First: "Paul"}

	PrintFirst(p)
	// Out: Paul
}
