package main

import (
	"fmt"
)

func main() {
	greeting("John")
}

func greeting(name string) {
	fmt.Printf("hello... %s\n", name)
}
