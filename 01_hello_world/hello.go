package main

import "fmt"

func Hello() string {
	return "Hello, world"
}

func Greet(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello())
}