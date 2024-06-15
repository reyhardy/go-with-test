package main

import (
	"fmt"
)

var name string

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello(name))
}
