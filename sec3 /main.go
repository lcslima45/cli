package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "anon", "name")
	flag.Parse()
	greeting(*name)
}

func greeting(name string) {
	fmt.Printf("hello... %s\n", name)
}