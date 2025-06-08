package golang

import (
	"testing"
)

func TestFormatNode(t *testing.T) {
	input := "package main; func  main (   ) { println(\"hello\") }"
	expected := `package main

func main() { println("hello") }
`
	output, err := FormatNode(input)
	if err != nil {
		t.Errorf("FormatNode returned an error: %v", err)
	}
	if output != expected {
		t.Errorf("FormatNode output mismatch:\nExpected:\n%s\nGot:\n%s", expected, output)
	}
}

func TestFormatNodeWithError(t *testing.T) {
	input := "package main; func main() { println(\"hello\"" // Missing closing parenthesis
	_, err := FormatNode(input)
	if err == nil {
		t.Errorf("FormatNode should have returned an error for invalid Go code")
	}
}
