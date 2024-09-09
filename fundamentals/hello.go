package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name
}

const englishHelloPrefix = "Hello, "

func CleverHello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("James"))
}
