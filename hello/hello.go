package main

import "fmt"

const french = "French"
const spanish = "Spanish"

const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "
const englishHelloPrefix = "Hello, "

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = englishHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// Hello does wat.
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("world", "English"))
}
