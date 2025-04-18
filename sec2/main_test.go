package main

import (
	"fmt"
	"io"
	"os"
	"testing"
)



func Test_greeting(t *testing.T) {
	originalStdout := os.Stdout
	r, w, err := os.Pipe()
	if  err != nil {
		t.Fatal(err)
	}
	os.Stdout = w
	name := "John"
	greeting(name)
	_ = w.Close()
	os.Stdout = originalStdout
	expeted := "hello... John\n"

	bs, err := io.ReadAll(r)

	if err != nil {
		t.Fatal(err)
	}

	got := string(bs)
	if expeted != got {
		t.Errorf("expected %q but got %q", expeted, got)
	}
	_ = r.Close()
	fmt.Println("Done.")
}
