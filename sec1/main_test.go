package main

import (
	"bytes"
	"testing"
)


func Test_greeting(t *testing.T) {
	buff := new(bytes.Buffer)
	name := "John"
	greeting(buff, name)
	expected := "hello... John\n"
	got := buff.String()
	if expected != got  {
		t.Errorf("expected %q but got %q", expected, got)
	}
}
