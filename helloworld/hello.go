package main

import (
	"fmt"
)

const englishHelloPerfix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "

const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return correctedLanguagePrefix(language) + name
}

func correctedLanguagePrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPerfix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
