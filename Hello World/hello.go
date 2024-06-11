package main

import "fmt"

const prefix = "Hello "

func hello(name string) string {
	if name == "" {
		name = "World"
	}
	return prefix + name
}

func main() {
	fmt.Println(hello("Anuj"))
}
