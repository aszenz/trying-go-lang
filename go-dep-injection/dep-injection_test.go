package main

import (
	"bytes"
	"testing"

	"github.com/aszenz/basic-golang/go-test-helpers"
)

func TestPrintGreeting(t *testing.T) {
	buffer := bytes.Buffer{}
	greeting := "Hello me"
	PrintGreeting(&buffer, greeting)
	got := buffer.String()
	want := greeting
	testhelpers.AssertStringsEqual(t, got, want)
}
