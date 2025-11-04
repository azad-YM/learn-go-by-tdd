package main

import "fmt"

const englishHelloPrifix = "Hello, "

func Hello(name string) string {
	if name == "" {
		return englishHelloPrifix + "World"
	}

	return englishHelloPrifix + name
}

func main() {
	fmt.Println(Hello(""))
}
