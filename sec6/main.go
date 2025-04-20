package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "", "username")
	password := flag.String("password", "", "password")
	flag.Parse() 
	var port int 
	flag.IntVar(&port, "port", 8080, "db port")

	fmt.Println("username", *username, "password", *password)
}