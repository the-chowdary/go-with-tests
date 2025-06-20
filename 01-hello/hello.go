package main

import (
	"fmt"
	"strings"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "world!"
	}
	// if language == "" || strings.ToLower(language) == "english" {
	// 	language = englishHelloPrefix
	// } else if strings.ToLower(language) == "spanish" {
	// 	language = spanishHelloPrefix
	// } else {
	// 	language = frenchHelloPrefix
	// }
	// return language + name

	prefix := englishHelloPrefix

	switch strings.ToLower(language) {
	case "spanish":
		prefix = spanishHelloPrefix
	case "french":
		prefix = frenchHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("The Chowdary", "Spanish"))
}
