package main

import "fmt"

const helloPrefix = "Hello, "

func hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(hello("world", ""))
}
