package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello() string {
	return "Hello, world"
}

func Greet(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello())
	fmt.Println(Greet("Chris"))
}
