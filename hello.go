package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const italian = "Italian"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const italianHelloPrefix = "Ciao, "

//exported with Capital letter in name. PUBLIC
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

//!= exported due to lowercase letter in name. This is because it is internal to algo
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case italian:
		prefix = italianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
